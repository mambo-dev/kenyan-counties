package internal

import (
	"time"

	"github.com/google/uuid"
)

type CountyParams struct {
	Name string `json:"name"`
}

type CountyResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	CountyID int64     `json:"county_id"`
}

type ApiResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type SubCountyResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	CountyID         string    `json:"county_id"`
	SubCountyGivenID int64     `json:"sub_county_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type WardResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	SubCountyID string    `json:"sub_county_id"`
	WardGivenID int64     `json:"ward_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaginatedResponse struct {
	TotalCount int64       `json:"total_count"`
	Data       interface{} `json:"items"`
}
