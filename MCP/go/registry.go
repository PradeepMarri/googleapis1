package main

import (
	"github.com/grafeas-proto/mcp-server/config"
	"github.com/grafeas-proto/mcp-server/models"
	tools_projects "github.com/grafeas-proto/mcp-server/tools/projects"
	tools_grafeas "github.com/grafeas-proto/mcp-server/tools/grafeas"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_projects.CreateProjects_listprojectsTool(cfg),
		tools_projects.CreateProjects_createprojectTool(cfg),
		tools_grafeas.CreateGrafeas_deleteoccurrence2Tool(cfg),
		tools_grafeas.CreateGrafeas_getoccurrence2Tool(cfg),
		tools_grafeas.CreateGrafeas_updateoccurrence2Tool(cfg),
		tools_grafeas.CreateGrafeas_listnoteoccurrences2Tool(cfg),
		tools_grafeas.CreateGrafeas_deletenoteTool(cfg),
		tools_grafeas.CreateGrafeas_getnoteTool(cfg),
		tools_grafeas.CreateGrafeas_updatenoteTool(cfg),
		tools_grafeas.CreateGrafeas_getoccurrencenote2Tool(cfg),
		tools_grafeas.CreateGrafeas_getnote2Tool(cfg),
		tools_grafeas.CreateGrafeas_updatenote2Tool(cfg),
		tools_grafeas.CreateGrafeas_deletenote2Tool(cfg),
		tools_grafeas.CreateGrafeas_deleteoccurrenceTool(cfg),
		tools_grafeas.CreateGrafeas_getoccurrenceTool(cfg),
		tools_grafeas.CreateGrafeas_updateoccurrenceTool(cfg),
		tools_grafeas.CreateGrafeas_batchcreateoccurrences2Tool(cfg),
		tools_projects.CreateProjects_deleteprojectTool(cfg),
		tools_projects.CreateProjects_getprojectTool(cfg),
		tools_grafeas.CreateGrafeas_batchcreatenotes2Tool(cfg),
		tools_grafeas.CreateGrafeas_batchcreatenotesTool(cfg),
		tools_grafeas.CreateGrafeas_listoccurrencesTool(cfg),
		tools_grafeas.CreateGrafeas_createoccurrenceTool(cfg),
		tools_grafeas.CreateGrafeas_listnotes2Tool(cfg),
		tools_grafeas.CreateGrafeas_createnote2Tool(cfg),
		tools_grafeas.CreateGrafeas_listnotesTool(cfg),
		tools_grafeas.CreateGrafeas_createnoteTool(cfg),
		tools_grafeas.CreateGrafeas_batchcreateoccurrencesTool(cfg),
		tools_grafeas.CreateGrafeas_getoccurrencenoteTool(cfg),
		tools_grafeas.CreateGrafeas_listoccurrences2Tool(cfg),
		tools_grafeas.CreateGrafeas_createoccurrence2Tool(cfg),
		tools_grafeas.CreateGrafeas_listnoteoccurrencesTool(cfg),
	}
}
