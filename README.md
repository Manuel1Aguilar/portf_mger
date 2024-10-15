# Portfolio Manager project

This project is based on the quote by Charlie Munger:
“If all you ever did was buy high-quality stocks on the 200-week moving average, you would beat the S&P 500 by a large margin over time.
The problem is, few human beings have that kind of discipline."

So the functionalities are:

- Search for stocks or crypto by symbol (Get current price and 200-week MA)
- Add asset to list of followed stocks (Stock + intended % of portfolio)
- Alert when an asset should be bought
- Save asset buys/sells. Keep record of the portfolio

Initially it will be a CLI. There's space to think about doing a web-page and mobile app from this.

## Developed commands:

### add-asset: Example usage "add-asset <symbol> <description> <type>"

Adds an asset to be tracked by the portfolio manager. This asset will be added to the list of assets for  which the asset snapshots will be generated, flags and targets will be checked and this asset can now be part of a transaction to be added to the portfolio. The supported asset types are: "CRYPTO" and "STOCK"

### get-asset: Example usage "get-asset <symbol>"

Gets the saved asset object (symbol, description and type) in the database if it exists

### get-assets: Example usae "get-assets"

Gets all the saved asset objects (symbol, description and type that exist in the database

### search-stock: Example usage "search-stock <symbol>"

Searches current stock value and calculates 200 weeks moving average for the stock of the corresponding symbol. Doesn't retrieve nor save anything from or to the database.

### set-objective: Example usage "set-objective <symbol> <target allocation %>"

Creates or updates the target allocation percentage for a given asset on the portfolio. E.g. if the current amount of NVDA stock represents 10% of the portfolio and the objective is 15% on the next update the NVDA stock is near it's 200 week moving average a flag to buy will be raised.

### transact: Example usage "transact <symbol> <type> <value in USD> <units bought>"

Creates a new transaction for the asset corresponding to the symbol. Type can be either BUY or SELL. This also updates the portfolio holding for this asset

### pfolio-status: Example usage "pfolio-status"

Gets the current portfolio status. It first updates all the holdings with the current value for the asset and calculates the percentages.
