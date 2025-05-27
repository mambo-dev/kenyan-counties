-- +goose Up
CREATE TABLE counties (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  county_given_id TEXT UNIQUE NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS counties;
