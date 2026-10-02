package main

import (
	"log"

	"Jarvis/auth"
	"Jarvis/brain"
	"Jarvis/clients"
	"Jarvis/config"
	"Jarvis/tools/app_launcher"
	"Jarvis/tools/registry"
	"Jarvis/tools/websearch"
)

func main() {
	cfg := config.Load()

	api := clients.NewAPIClient(cfg.APIURL)

	username, password, err := auth.GetCredentials()
	if err != nil {
		log.Fatal("Could not get credentials:", err)
	}

	if err := api.Login(username, password); err != nil {
		log.Fatal("Login failed:", err)
	}
	searcher := websearch.NewSearXNGSearcher(cfg.SearXNGURL)
	toolRegistry := registry.New()

	webSearchTool := websearch.NewTool(searcher)

	toolRegistry.Register(webSearchTool)

	appLauncher := applauncher.NewTool()
	toolRegistry.Register(appLauncher)
	jarvis := brain.NewBrain(api, toolRegistry)

	if err := jarvis.StartConversation(); err != nil {
		log.Fatal("Could not start conversation:", err)
	}

	if err := jarvis.StartSession(); err != nil {
		log.Fatal("Session error:", err)
	}
}
