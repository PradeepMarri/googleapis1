package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/grafeas-proto/mcp-server/config"
	"github.com/grafeas-proto/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Grafeas_listnoteoccurrences2Handler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["filter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter=%v", val))
		}
		if val, ok := args["pageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageSize=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/%s/occurrences%s", cfg.BaseURL, name_1, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.V1ListNoteOccurrencesResponse
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

func CreateGrafeas_listnoteoccurrences2Tool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_name_1_occurrences",
		mcp.WithDescription("Lists occurrences referencing the specified note. Provider projects can use
this method to get all occurrences across consumer projects referencing the
specified note."),
		mcp.WithString("name_1", mcp.Required(), mcp.Description("The name of the note to list occurrences for in the form of\n`projects/[PROVIDER_ID]/notes/[NOTE_ID]`.")),
		mcp.WithString("filter", mcp.Description("The filter expression.")),
		mcp.WithNumber("pageSize", mcp.Description("Number of occurrences to return in the list.")),
		mcp.WithString("pageToken", mcp.Description("Token to provide to skip to a particular spot in the list.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Grafeas_listnoteoccurrences2Handler(cfg),
	}
}
