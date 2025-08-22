package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/grafeas-proto/mcp-server/config"
	"github.com/grafeas-proto/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Grafeas_createnoteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		parentVal, ok := args["parent"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: parent"), nil
		}
		parent, ok := parentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: parent"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["noteId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("noteId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.V1Note
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v1/%s/notes%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.V1Note
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGrafeas_createnoteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_parent_notes",
		mcp.WithDescription("Creates a new note."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("The name of the project in the form of `projects/[PROJECT_ID]`, under which\nthe note is to be created.")),
		mcp.WithString("noteId", mcp.Required(), mcp.Description("The ID to use for this note.")),
		mcp.WithObject("build", mcp.Description("Input parameter: Note holding the version of the provider's builder and the signature of the\nprovenance message in the build details occurrence.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Output only. The time this note was last updated. This field can be used as\na filter in list requests.")),
		mcp.WithArray("relatedNoteNames", mcp.Description("Input parameter: Other notes related to this note.")),
		mcp.WithString("longDescription", mcp.Description("Input parameter: A detailed description of this note.")),
		mcp.WithString("shortDescription", mcp.Description("Input parameter: A one sentence description of this note.")),
		mcp.WithObject("upgrade", mcp.Description("Input parameter: An Upgrade Note represents a potential upgrade of a package to a given\nversion. For each package version combination (i.e. bash 4.0, bash 4.1,\nbash 4.1.2), there will be an Upgrade Note. For Windows, windows_update field\nrepresents the information related to the update.")),
		mcp.WithObject("image", mcp.Description("Input parameter: Basis describes the base image portion (Note) of the DockerImage\nrelationship. Linked occurrences are derived from this or an equivalent image\nvia:\n  FROM <Basis.resource_url>\nOr an equivalent reference, e.g., a tag of the resource_url.")),
		mcp.WithString("name", mcp.Description("Input parameter: Output only. The name of the note in the form of\n`projects/[PROVIDER_ID]/notes/[NOTE_ID]`.")),
		mcp.WithObject("deployment", mcp.Description("Input parameter: An artifact that can be deployed in some runtime.")),
		mcp.WithObject("attestation", mcp.Description("Input parameter: Note kind that represents a logical attestation \"role\" or \"authority\". For\nexample, an organization might have one `Authority` for \"QA\" and one for\n\"build\". This note is intended to act strictly as a grouping mechanism for\nthe attached occurrences (Attestations). This grouping mechanism also\nprovides a security boundary, since IAM ACLs gate the ability for a principle\nto attach an occurrence to a given note. It also provides a single point of\nlookup to find all attached attestation occurrences, even if they don't all\nlive in the same project.")),
		mcp.WithObject("discovery", mcp.Description("Input parameter: A note that indicates a type of analysis a provider would perform. This note\nexists in a provider's project. A `Discovery` occurrence is created in a\nconsumer's project at the start of analysis.")),
		mcp.WithString("expirationTime", mcp.Description("Input parameter: Time of expiration for this note. Empty if note does not expire.")),
		mcp.WithArray("relatedUrl", mcp.Description("Input parameter: URLs associated with this note.")),
		mcp.WithObject("compliance", mcp.Description("")),
		mcp.WithObject("package", mcp.Description("Input parameter: PackageNote represents a particular package version.")),
		mcp.WithObject("secret", mcp.Description("Input parameter: The note representing a secret.")),
		mcp.WithObject("dsseAttestation", mcp.Description("")),
		mcp.WithObject("vulnerability", mcp.Description("Input parameter: A security vulnerability that can be found in resources.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Kind represents the kinds of notes supported.\n\n - NOTE_KIND_UNSPECIFIED: Default value. This value is unused.\n - VULNERABILITY: The note and occurrence represent a package vulnerability.\n - BUILD: The note and occurrence assert build provenance.\n - IMAGE: This represents an image basis relationship.\n - PACKAGE: This represents a package installed via a package manager.\n - DEPLOYMENT: The note and occurrence track deployment events.\n - DISCOVERY: The note and occurrence track the initial discovery status of a resource.\n - ATTESTATION: This represents a logical \"role\" that can attest to artifacts.\n - UPGRADE: This represents an available package upgrade.\n - COMPLIANCE: This represents a Compliance Note\n - DSSE_ATTESTATION: This represents a DSSE attestation Note\n - VULNERABILITY_ASSESSMENT: This represents a Vulnerability Assessment.\n - SBOM_REFERENCE: This represents an SBOM Reference.\n - SECRET: This represents a secret.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. The time this note was created. This field can be used as a\nfilter in list requests.")),
		mcp.WithObject("sbomReference", mcp.Description("Input parameter: The note representing an SBOM reference.")),
		mcp.WithObject("vulnerabilityAssessment", mcp.Description("Input parameter: A single VulnerabilityAssessmentNote represents\none particular product's vulnerability assessment for one CVE.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Grafeas_createnoteHandler(cfg),
	}
}
