package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/grafeas-proto/mcp-server/config"
	"github.com/grafeas-proto/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Grafeas_deleteoccurrence2Handler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/v1/%s", cfg.BaseURL, name_1)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result interface{}
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

func CreateGrafeas_deleteoccurrence2Tool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_v1_name_1",
		mcp.WithDescription("Deletes the specified occurrence. For example, use this method to delete an
occurrence when the occurrence is no longer applicable for the given
resource."),
		mcp.WithString("name_1", mcp.Required(), mcp.Description("The name of the occurrence in the form of\n`projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Grafeas_deleteoccurrence2Handler(cfg),
	}
}
