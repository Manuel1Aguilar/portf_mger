-- Add asset_objective table
CREATE TABLE IF NOT EXISTS asset_objective (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    asset_id INTEGER NOT NULL,
    target_allocation_percentage REAL NOT NULL,
    FOREIGN KEY (asset_id) REFERENCES asset(id)
);
