-- Add portfolio_entry table
CREATE TABLE IF NOT EXISTS portfolio_entry (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    asset_id INTEGER NOT NULL,
    units_held REAL NOT NULL,
    current_position REAL NOT NULL,
    FOREIGN KEY (asset_id) REFERENCES asset(id)
);
