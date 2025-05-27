-- name: GetCountyByName :one
SELECT * FROM counties
WHERE name = $1;

-- name: SearchCountiesByName :many
SELECT * FROM counties
WHERE name ILIKE '%' || $1 || '%'
ORDER BY name
LIMIT $2 OFFSET $3;

-- name: CreateCounty :one
INSERT INTO counties (name, county_given_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListCounties :many
SELECT * FROM counties
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: DeleteCountyByID :exec
DELETE FROM counties
WHERE id = $1;
