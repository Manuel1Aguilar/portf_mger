-- Add portfolio_holding table
CREATE TABLE IF NOT EXISTS portfolio_holding (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    asset_id INTEGER NOT NULL,
    units_held REAL NOT NULL,
    usd_value REAL NOT NULL,
    last_updated TEXT NOT NULL,
    FOREIGN KEY (asset_id) REFERENCES asset(id)
);
