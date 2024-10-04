package main

import (
	"log"

	"github.com/Manuel1Aguilar/portf_mger/internal/app"
	"github.com/Manuel1Aguilar/portf_mger/internal/commands"
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
