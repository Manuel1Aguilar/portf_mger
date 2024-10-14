-- Add asset_transaction table
CREATE TABLE IF NOT EXISTS asset_transaction(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    asset_id INTEGER NOT NULL,
    transaction_type TEXT NOT NULL, -- BUY or SELL
    valueUSD REAL NOT NULL,
    units REAL NOT NULL,
    unit_price REAL NOT NULL,
    date_transacted DATETIME NOT NULL,
    FOREIGN KEY (asset_id) REFERENCES asset(id)
);
