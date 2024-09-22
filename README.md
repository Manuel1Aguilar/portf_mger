# Stock Tracker project

This project is based on the quote by Charlie Munger:
“If all you ever did was buy high-quality stocks on the 200-week moving average, you would beat the S&P 500 by a large margin over time. The problem is, few human beings have that kind of discipline."

So the functionalities are:

- Search for stocks by symbol (Get current price and 200-week MA)
- Add stock to list of followed stocks (Stock + intended % of portfolio)
- Alert when a stock should be bought
- Save stock buys/sells. Keep record of the portfolio

Initially it will be a CLI. There's space to think about doing a web-page and mobile app from this.

## Project To-Do List

### 1. POC CLI Interaction
- [x] Set up the project folder structure.
- [x] Create a basic command-line interface (CLI) to interact with the app.
- [x] Use the CLI to prompt user input for searching and displaying stock data.
- [x] Output results in the terminal.

### 2. POC Stock API Interaction
- [ ] Research and select a stock market API (e.g., Alpha Vantage, Yahoo Finance).
- [ ] Write a function to fetch stock data from the selected API.
- [ ] Parse the response and extract relevant data (e.g., historical stock prices).

### 3. Search Stocks
- [ ] Implement a function to search for a specific stock by its symbol.
- [ ] Display the stock’s current data (price, volume, etc.) in the terminal.
- [ ] Save the stock data locally using SQLite.

### 4. Calculate 200 MA
- [ ] Write a function to calculate the 200-week moving average (MA) for a stock.
- [ ] Display the 200 MA value for a searched stock.
- [ ] Flag stocks that are near or below their 200 MA for buying opportunities.

### 5. Add Tasks
- [ ] Implement portfolio management logic to track investments and percentage allocation per stock.
- [ ] Allow the user to assign objective percentages for each stock in the portfolio.
- [ ] Write a rebalancing function that checks if the current allocation matches the objective.
- [ ] Store portfolio data (current and objective percentages) in SQLite.
- [ ] Display portfolio balance and flags for rebalancing.

---

### Future Improvements:
- [ ] Add GUI interface options (after completing the terminal interface).
- [ ] Implement more advanced stock analysis features (like other moving averages or indicators).
- [ ] Consider integrating a more robust database (PostgreSQL) as the app scales.


