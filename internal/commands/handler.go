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
	case "get-assets":
		err := getAssets(application)
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
	case "set-objective":
		err := setObjective(application)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "transact":
		err := transact(application)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "pfolio-status":
		err := getPortfolioStatus(application)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "help":
		help()
	default:
		fmt.Println("Command not found")
		help()
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

func getAssets(application *app.App) error {
	assets, err := application.AssetService.GetAssets()
	if err != nil {
		return err
	}
	fmt.Println("Assets:")
	for _, asset := range assets {
		fmt.Printf("%v \n", asset)
	}
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

func setObjective(application *app.App) error {
	if len(os.Args) < 4 {
		return fmt.Errorf("Example usage: set-objective <symbol> <target allocation %%>")
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

func transact(application *app.App) error {
	if len(os.Args) < 5 {
		return fmt.Errorf("Example usage: transact <symbol> <type> <Units bought>")
	}
	symbol := os.Args[2]
	transType := os.Args[3]
	if transType != "BUY" && transType != "SELL" {
		return fmt.Errorf("Type must be either BUY or SELL")
	}

	units, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		return fmt.Errorf("Error parsing value in USD from input: %v", err)
	}

	createModel := &models.AssetTransactionCreate{
		Symbol: symbol,
		Type:   transType,
		Units:  units,
	}
	err = application.AssetTransactionService.SaveAssetTransaction(createModel)
	if err != nil {
		return fmt.Errorf("Error saving transaction: %v", err)
	}

	fmt.Println("Transaction created")
	return nil
}

func getPortfolioStatus(application *app.App) error {
	pfolio, err := application.PortfolioHoldingService.GetUpdatedPortfolio()
	if err != nil {
		fmt.Printf("%v \n", err)
		return fmt.Errorf("Error getting the updated portfolio")
	}
	fmt.Printf(pfolio.String() + "\n")
	return nil
}

func help() {
	fmt.Println("Available Commands:")
	fmt.Println("add-asset <symbol> <description> <type> ; E.g.: add-asset NVDA Nvidia STOCK; Types: STOCK/CRYPTO")
	fmt.Println("get-asset <symbol>")
	fmt.Println("get-assets")
	fmt.Println("search-stock <symbol>")
	fmt.Println("set-objective <symbol> <target allocation %>")
	fmt.Println("transact <symbol> <type> <units bought>")
	fmt.Println("pfolio-status")
}
