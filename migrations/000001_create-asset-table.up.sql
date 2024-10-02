-- Add asset table
CREATE TABLE IF NOT EXISTS asset (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    symbol TEXT NOT NULL,
    description TEXT,
    asset_type TEXT NOT NULL -- CRYPTO or STOCK
);
