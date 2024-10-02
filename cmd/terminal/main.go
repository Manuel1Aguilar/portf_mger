package main

import (
	"log"

	"stock_tracker/internal/app"
	"stock_tracker/internal/commands"
)

func main() {
	// Initialize the app
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize the app: %v", err)
	}
	defer application.Close()

	commands.HandleCommand(application)
}
