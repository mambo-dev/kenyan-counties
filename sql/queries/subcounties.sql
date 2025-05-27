-- name: GetSubCountyByID :one
SELECT * FROM sub_counties
WHERE id = ?;

-- name: GetSubCountyByName :one
SELECT * FROM sub_counties
WHERE name = ?;

-- name: GetSubCountyByGivenID :one
SELECT * FROM sub_counties
WHERE sub_county_given_id = ?;

-- name: GetSubCountiesByCountyID :many
SELECT * FROM sub_counties
WHERE county_id = ?
ORDER BY sub_county_given_id
LIMIT ? OFFSET ?;

-- name: SearchSubCountiesByName :many
SELECT * FROM sub_counties
WHERE LOWER(name) LIKE '%' || LOWER(?) || '%'
ORDER BY sub_county_given_id
LIMIT ? OFFSET ?;

-- name: CreateSubCounty :one
INSERT INTO sub_counties (id, name, county_id, sub_county_given_id)
VALUES (?, ?, ?, ?)
RETURNING *;



-- name: ListSubCounties :many
SELECT * FROM sub_counties
ORDER BY sub_county_given_id
LIMIT ? OFFSET ?;

-- name: TotalSubCounties :one
SELECT COUNT(*) AS total FROM sub_counties;
