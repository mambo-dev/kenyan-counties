-- +goose Up
CREATE TABLE sub_counties (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    county_id UUID NOT NULL REFERENCES counties(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_sub_counties_county_id ON sub_counties(county_id);

-- +goose Down
DROP INDEX IF EXISTS idx_sub_counties_county_id;
DROP TABLE IF EXISTS sub_counties;
