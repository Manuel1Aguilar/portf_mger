package commands

import (
	"fmt"
	"os"

	"github.com/Manuel1Aguilar/portf_mger/internal/api"
	"github.com/Manuel1Aguilar/portf_mger/internal/app"
	"github.com/Manuel1Aguilar/portf_mger/internal/models"
)

func HandleCommand(application *app.App) {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]

	switch command {
	case "add-asset":
		if len(os.Args) < 3 {
			fmt.Println("Example usage: add-asset <symbol> <description> <type> ; Types: CRYPTO / STOCK")
			return
		}

		symbol := os.Args[2]
		description := os.Args[3]
		aType := os.Args[4]

		if aType != "CRYPTO" && aType != "STOCK" {
			fmt.Printf("Asset type has to be either CRYPTO or STOCK. Provided: %s \n", aType)
			return
		}

		asset := &models.Asset{
			Symbol:      symbol,
			Description: description,
			AssetType:   aType,
		}

		err := application.AssetService.CreateAsset(asset)
		if err != nil {
			fmt.Printf("Error creating new asset: %v\n", err)
			return
		}

		fmt.Printf("Succesfully added asset: %s \n", symbol)

		fmt.Println(asset)
	case "get-asset":
		if len(os.Args) < 3 {
			fmt.Println("Example usage: get-asset <symbol>")
			return
		}
		symbol := os.Args[2]
		asset, err := application.AssetService.GetAssetBySymbol(symbol)
		if err != nil {
			fmt.Printf("Error executing get-asset command: %v", err)
		}
		fmt.Printf("Asset for %s:\n", symbol)
		fmt.Println(asset)
	case "search-stock":
		if len(os.Args) < 3 {
			fmt.Println("Example usage: search-stock <symbol>")
			return
		}

		symbol := os.Args[2]
		ma, err := api.Get200WeekMovingAverage(symbol)
		if err != nil {
			fmt.Printf("Error getting 200w MA: %v\n", err)
		}
		fmt.Printf("%v\n", ma)
	default:
		fmt.Println("Command not found")
	}

}
