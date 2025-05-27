-- +goose Up
CREATE TABLE wards (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  sub_county_id TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (sub_county_id) REFERENCES sub_counties(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS wards;
