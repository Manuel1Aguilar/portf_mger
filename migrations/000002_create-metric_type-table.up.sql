-- Add metric_type table
CREATE TABLE IF NOT EXISTS metrict_type (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type_name TEXT NOT NULL, -- Example: 'ma_200w', 'ma_100w'
    description TEXT -- Description of the metric
);
