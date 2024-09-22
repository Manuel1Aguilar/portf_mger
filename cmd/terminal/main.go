package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"stock_tracker/internal/api"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
	apikey := os.Getenv("API_KEY")
	fmt.Println("API_KEY", apikey)

	fmt.Println("Welcome to the Stock Tracker CLI")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter a command (search [symbol], help, exit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch {
		case input == "exit":
			fmt.Println("Goodbye!")
			return
		case input == "help":
			displayHelp()
		case strings.HasPrefix(input, "search"):
			symbol := strings.Trim(strings.TrimPrefix(input, "search"), " ")
			searchStock(symbol)
		default:
			fmt.Println("Invalid command. Type 'help' for a list of commmands.")
		}
	}
}

func displayHelp() {
	fmt.Println("Available commands:")
	fmt.Println("Help: Display help")
	fmt.Println("Search [symbol]: Search for a stock. Example: search AAPL")
	fmt.Println("Exit: Exit the program")
}

func searchStock(symbol string) {
	fmt.Printf("Searching for stock: %s...\n", symbol)
	data, err := api.FetchStockData(symbol)
	if err != nil {
		fmt.Printf("Error while fetching stock %s data: %s", symbol, err)
	}

	fmt.Println(data)
}
