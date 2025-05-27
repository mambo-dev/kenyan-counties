-- +goose Up
CREATE TABLE wards (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    sub_county_id UUID NOT NULL REFERENCES sub_counties(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
CREATE INDEX idx_wards_sub_county_id ON wards(sub_county_id);
-- +goose Down
DROP INDEX IF EXISTS idx_wards_sub_county_id;
DROP TABLE IF EXISTS wards;
