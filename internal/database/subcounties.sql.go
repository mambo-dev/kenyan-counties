// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: subcounties.sql

package database

import (
	"context"
)

const createSubCounty = `-- name: CreateSubCounty :one
INSERT INTO sub_counties (id, name, county_id, sub_county_given_id)
VALUES (?, ?, ?, ?)
RETURNING id, name, county_id, sub_county_given_id, created_at, updated_at
`

type CreateSubCountyParams struct {
	ID               string
	Name             string
	CountyID         string
	SubCountyGivenID int64
}

func (q *Queries) CreateSubCounty(ctx context.Context, arg CreateSubCountyParams) (SubCounty, error) {
	row := q.db.QueryRowContext(ctx, createSubCounty,
		arg.ID,
		arg.Name,
		arg.CountyID,
		arg.SubCountyGivenID,
	)
	var i SubCounty
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CountyID,
		&i.SubCountyGivenID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSubCountiesByCountyID = `-- name: GetSubCountiesByCountyID :many
SELECT id, name, county_id, sub_county_given_id, created_at, updated_at FROM sub_counties
WHERE county_id = ?
ORDER BY sub_county_given_id
LIMIT ? OFFSET ?
`

type GetSubCountiesByCountyIDParams struct {
	CountyID string
	Limit    int64
	Offset   int64
}

func (q *Queries) GetSubCountiesByCountyID(ctx context.Context, arg GetSubCountiesByCountyIDParams) ([]SubCounty, error) {
	rows, err := q.db.QueryContext(ctx, getSubCountiesByCountyID, arg.CountyID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SubCounty
	for rows.Next() {
		var i SubCounty
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CountyID,
			&i.SubCountyGivenID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubCountyByGivenID = `-- name: GetSubCountyByGivenID :one
SELECT id, name, county_id, sub_county_given_id, created_at, updated_at FROM sub_counties
WHERE sub_county_given_id = ?
`

func (q *Queries) GetSubCountyByGivenID(ctx context.Context, subCountyGivenID int64) (SubCounty, error) {
	row := q.db.QueryRowContext(ctx, getSubCountyByGivenID, subCountyGivenID)
	var i SubCounty
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CountyID,
		&i.SubCountyGivenID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSubCountyByID = `-- name: GetSubCountyByID :one
SELECT id, name, county_id, sub_county_given_id, created_at, updated_at FROM sub_counties
WHERE id = ?
`

func (q *Queries) GetSubCountyByID(ctx context.Context, id string) (SubCounty, error) {
	row := q.db.QueryRowContext(ctx, getSubCountyByID, id)
	var i SubCounty
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CountyID,
		&i.SubCountyGivenID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSubCountyByName = `-- name: GetSubCountyByName :one
SELECT id, name, county_id, sub_county_given_id, created_at, updated_at FROM sub_counties
WHERE name = ?
`

func (q *Queries) GetSubCountyByName(ctx context.Context, name string) (SubCounty, error) {
	row := q.db.QueryRowContext(ctx, getSubCountyByName, name)
	var i SubCounty
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CountyID,
		&i.SubCountyGivenID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listSubCounties = `-- name: ListSubCounties :many
SELECT id, name, county_id, sub_county_given_id, created_at, updated_at FROM sub_counties
ORDER BY sub_county_given_id
LIMIT ? OFFSET ?
`

type ListSubCountiesParams struct {
	Limit  int64
	Offset int64
}

func (q *Queries) ListSubCounties(ctx context.Context, arg ListSubCountiesParams) ([]SubCounty, error) {
	rows, err := q.db.QueryContext(ctx, listSubCounties, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SubCounty
	for rows.Next() {
		var i SubCounty
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CountyID,
			&i.SubCountyGivenID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchSubCountiesByName = `-- name: SearchSubCountiesByName :many
SELECT id, name, county_id, sub_county_given_id, created_at, updated_at FROM sub_counties
WHERE LOWER(name) LIKE '%' || LOWER(?) || '%'
ORDER BY sub_county_given_id
LIMIT ? OFFSET ?
`

type SearchSubCountiesByNameParams struct {
	LOWER  string
	Limit  int64
	Offset int64
}

func (q *Queries) SearchSubCountiesByName(ctx context.Context, arg SearchSubCountiesByNameParams) ([]SubCounty, error) {
	rows, err := q.db.QueryContext(ctx, searchSubCountiesByName, arg.LOWER, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SubCounty
	for rows.Next() {
		var i SubCounty
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CountyID,
			&i.SubCountyGivenID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const totalSubCounties = `-- name: TotalSubCounties :one
SELECT COUNT(*) AS total FROM sub_counties
`

func (q *Queries) TotalSubCounties(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, totalSubCounties)
	var total int64
	err := row.Scan(&total)
	return total, err
}
