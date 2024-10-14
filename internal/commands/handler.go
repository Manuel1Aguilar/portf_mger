package commands

import (
	"fmt"
	"os"
	"strconv"

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
		err := addAsset(application)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "get-asset":
		err := getAsset(application)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "search-stock":
		err := searchStock()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "new-objective":
		err := newObjective(application)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "transact":
		err := newTransaction(application)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "pfolio-status":
		// TODO
		// Refresh all values
		// Check flags
		// Show a list of assets I hold w their value and % they represent with the intended %
	case "raise-flags":
		// TODO
		// Refresh all values
		// Check flags
		// Show a list of assets I should buy with their flags
	default:
		fmt.Println("Command not found")
	}

}

func addAsset(application *app.App) error {
	if len(os.Args) < 3 {
		return fmt.Errorf("Example usage: add-asset <symbol> <description> <type> ; Types: CRYPTO / STOCK")
	}

	symbol := os.Args[2]
	description := os.Args[3]
	aType := os.Args[4]

	if aType != "CRYPTO" && aType != "STOCK" {
		return fmt.Errorf("Asset type has to be either CRYPTO or STOCK. Provided: %s \n", aType)
	}

	asset := &models.Asset{
		Symbol:      symbol,
		Description: description,
		AssetType:   aType,
	}

	err := application.AssetService.CreateAsset(asset)
	if err != nil {
		return fmt.Errorf("Error creating new asset: %v\n", err)
	}

	fmt.Printf("Succesfully added asset: %s \n", symbol)

	fmt.Println(asset)

	return nil
}
func getAsset(application *app.App) error {
	if len(os.Args) < 3 {
		return fmt.Errorf("Example usage: get-asset <symbol>")
	}
	symbol := os.Args[2]
	asset, err := application.AssetService.GetAssetBySymbol(symbol)
	if err != nil {
		return fmt.Errorf("Error executing get-asset command: %v", err)
	}
	fmt.Printf("Asset for %s:\n", symbol)
	fmt.Println(asset)
	return nil
}

func searchStock() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("Example usage: search-stock <symbol>")
	}

	symbol := os.Args[2]
	ma, err := api.Get200WeekMovingAverage(symbol)
	if err != nil {
		return fmt.Errorf("Error getting 200w MA: %v\n", err)
	}
	fmt.Printf("%v\n", ma)
	return nil
}

func newObjective(application *app.App) error {
	if len(os.Args) < 4 {
		return fmt.Errorf("Example usage: new-objective <symbol> <target allocation %%>")
	}
	symbol := os.Args[2]
	taoc := os.Args[3]

	taocVal, err := strconv.ParseFloat(taoc, 64)
	if err != nil {
		return fmt.Errorf("Error parsing target allocation percentage (It has to be a number): %v", err)
	}
	createModel := &models.AssetObjectiveCreate{
		Symbol:                     symbol,
		TargetAllocationPercentage: taocVal,
	}
	err = application.PortfolioHoldingService.CreateAssetObjective(createModel)
	if err != nil {
		return fmt.Errorf("Error creating asset objective: %v", err)
	}
	return nil
}

func newTransaction(application *app.App) error {
	if len(os.Args) < 6 {
		return fmt.Errorf("Example usage: transact <symbol> <type> <Value in USD> <Units bought>")
	}
	symbol := os.Args[2]
	transType := os.Args[3]
	if transType != "BUY" && transType != "SELL" {
		return fmt.Errorf("Type must be either BUY or SELL")
	}
	valueUsd, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		return fmt.Errorf("Error parsing value in USD from input: %v", err)
	}

	units, err := strconv.ParseFloat(os.Args[5], 64)
	if err != nil {
		return fmt.Errorf("Error parsing value in USD from input: %v", err)
	}

	unitPrice := valueUsd / units
	createModel := &models.AssetTransactionCreate{
		Symbol:    symbol,
		Type:      transType,
		ValueUSD:  valueUsd,
		Units:     units,
		UnitPrice: unitPrice,
	}
	err = application.AssetTransactionService.SaveAssetTransaction(createModel)
	if err != nil {
		return fmt.Errorf("Error saving transaction: %v", err)
	}

	return nil
}
