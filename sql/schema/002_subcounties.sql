-- +goose Up
CREATE TABLE sub_counties (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  county_id TEXT NOT NULL,
  sub_county_given_id INT UNIQUE NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (county_id) REFERENCES counties(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS sub_counties;
