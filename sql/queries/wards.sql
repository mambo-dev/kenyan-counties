-- name: GetWardByID :one
SELECT * FROM wards
WHERE id = ?;

-- name: GetWardByGivenID :one
SELECT * FROM wards
WHERE ward_given_id = ?;

-- name: GetWardsBySubCountyID :many
SELECT * FROM wards
WHERE sub_county_id = ?
ORDER BY name
LIMIT ? OFFSET ?;

-- name: SearchWardsByName :many
SELECT * FROM wards
WHERE LOWER(name) LIKE '%' || LOWER(?) || '%'
ORDER BY name
LIMIT ? OFFSET ?;

-- name: CreateWard :one
INSERT INTO wards (id, name, sub_county_id, ward_given_id)
VALUES (?, ?, ?, ?)
RETURNING *;

