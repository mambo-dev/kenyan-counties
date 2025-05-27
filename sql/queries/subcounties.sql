-- name: GetSubCountyByID :one
SELECT * FROM sub_counties
WHERE id = ?;

-- name: GetSubCountiesByCountyID :many
SELECT * FROM sub_counties
WHERE county_id = ?
ORDER BY name
LIMIT ? OFFSET ?;

-- name: SearchSubCountiesByName :many
SELECT * FROM sub_counties
WHERE LOWER(name) LIKE '%' || LOWER(?) || '%'
ORDER BY name
LIMIT ? OFFSET ?;

-- name: CreateSubCounty :one
INSERT INTO sub_counties (id, name, county_id)
VALUES (?, ?, ?)
RETURNING *;

