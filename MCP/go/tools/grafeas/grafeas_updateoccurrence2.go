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

func Grafeas_updateoccurrence2Handler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		name_1Val, ok := args["name_1"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: name_1"), nil
		}
		name_1, ok := name_1Val.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: name_1"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["updateMask"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updateMask=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.V1Occurrence
		
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
		url := fmt.Sprintf("%s/v1/%s%s", cfg.BaseURL, name_1, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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
		var result models.V1Occurrence
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

func CreateGrafeas_updateoccurrence2Tool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_v1_name_1",
		mcp.WithDescription("Updates the specified occurrence."),
		mcp.WithString("name_1", mcp.Required(), mcp.Description("The name of the occurrence in the form of\n`projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.")),
		mcp.WithString("updateMask", mcp.Description("The fields to update.")),
		mcp.WithString("remediation", mcp.Description("Input parameter: A description of actions that can be taken to remedy the note.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Output only. The time this occurrence was last updated.")),
		mcp.WithObject("vulnerability", mcp.Description("Input parameter: An occurrence of a severity vulnerability on a resource.")),
		mcp.WithObject("compliance", mcp.Description("Input parameter: An indication that the compliance checks in the associated ComplianceNote\nwere not satisfied for particular resources or a specified reason.")),
		mcp.WithString("resourceUri", mcp.Description("Input parameter: Required. Immutable. A URI that represents the resource for which the\noccurrence applies. For example,\n`https://gcr.io/project/image@sha256:123abc` for a Docker image.")),
		mcp.WithObject("attestation", mcp.Description("Input parameter: Occurrence that represents a single \"attestation\". The authenticity of an\nattestation can be verified using the attached signature. If the verifier\ntrusts the public key of the signer, then verifying the signature is\nsufficient to establish trust. In this circumstance, the authority to which\nthis attestation is attached is primarily useful for lookup (how to find\nthis attestation if you already know the authority and artifact to be\nverified) and intent (for which authority this attestation was intended to\nsign.")),
		mcp.WithObject("build", mcp.Description("Input parameter: Details of a build occurrence.")),
		mcp.WithObject("discovery", mcp.Description("Input parameter: Provides information about the analysis status of a discovered resource.")),
		mcp.WithObject("envelope", mcp.Description("Input parameter: MUST match https://github.com/secure-systems-lab/dsse/blob/master/envelope.proto.\nAn authenticated message of arbitrary type.")),
		mcp.WithObject("secret", mcp.Description("Input parameter: The occurrence provides details of a secret.")),
		mcp.WithObject("dsseAttestation", mcp.Description("Input parameter: Deprecated. Prefer to use a regular Occurrence, and populate the\nEnvelope at the top level of the Occurrence.")),
		mcp.WithObject("image", mcp.Description("Input parameter: Details of the derived image portion of the DockerImage relationship. This\nimage would be produced from a Dockerfile with FROM <DockerImage.Basis in\nattached Note>.")),
		mcp.WithObject("upgrade", mcp.Description("Input parameter: An Upgrade Occurrence represents that a specific resource_url could install a\nspecific upgrade. This presence is supplied via local sources (i.e. it is\npresent in the mirror and the running system has noticed its availability).\nFor Windows, both distribution and windows_update contain information for the\nWindows update.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Kind represents the kinds of notes supported.\n\n - NOTE_KIND_UNSPECIFIED: Default value. This value is unused.\n - VULNERABILITY: The note and occurrence represent a package vulnerability.\n - BUILD: The note and occurrence assert build provenance.\n - IMAGE: This represents an image basis relationship.\n - PACKAGE: This represents a package installed via a package manager.\n - DEPLOYMENT: The note and occurrence track deployment events.\n - DISCOVERY: The note and occurrence track the initial discovery status of a resource.\n - ATTESTATION: This represents a logical \"role\" that can attest to artifacts.\n - UPGRADE: This represents an available package upgrade.\n - COMPLIANCE: This represents a Compliance Note\n - DSSE_ATTESTATION: This represents a DSSE attestation Note\n - VULNERABILITY_ASSESSMENT: This represents a Vulnerability Assessment.\n - SBOM_REFERENCE: This represents an SBOM Reference.\n - SECRET: This represents a secret.")),
		mcp.WithObject("package", mcp.Description("Input parameter: Details on how a particular software package was installed on a system.")),
		mcp.WithObject("sbomReference", mcp.Description("Input parameter: The occurrence representing an SBOM reference as applied to a specific\nresource. The occurrence follows the DSSE specification. See\nhttps://github.com/secure-systems-lab/dsse/blob/master/envelope.md for more\ndetails.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. The time this occurrence was created.")),
		mcp.WithString("name", mcp.Description("Input parameter: Output only. The name of the occurrence in the form of\n`projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.")),
		mcp.WithObject("deployment", mcp.Description("Input parameter: The period during which some deployable was active in a runtime.")),
		mcp.WithString("noteName", mcp.Description("Input parameter: Required. Immutable. The analysis note associated with this occurrence, in\nthe form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be\nused as a filter in list requests.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Grafeas_updateoccurrence2Handler(cfg),
	}
}
