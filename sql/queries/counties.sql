-- name: GetCountyByName :one
SELECT * FROM counties
WHERE name = ?;

-- name: SearchCountiesByName :many
SELECT * FROM counties
WHERE LOWER(name) LIKE '%' || LOWER(?) || '%'
ORDER BY name
LIMIT ? OFFSET ?;

-- name: CreateCounty :one
INSERT INTO counties (id, name, county_given_id)
VALUES (?, ?, ?)
RETURNING *;

-- name: ListCounties :many
SELECT * FROM counties
ORDER BY name
LIMIT ? OFFSET ?;

-- name: DeleteCountyByID :exec
DELETE FROM counties
WHERE id = ?;

