-- Add asset_snapshot table
CREATE TABLE IF NOT EXISTS asset_snapshot(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    asset_id INTEGER NOT NULL,
    metric_value REAL NOT NULL,
    metric_type_id INTEGER NOT NULL,
    current_value REAL NOT NULL,
    date_taken DATETIME NOT NULL,
    FOREIGN KEY (asset_id) REFERENCES asset(id),
    FOREIGN KEY (metric_type_id) REFERENCES metric_type(id)
);
