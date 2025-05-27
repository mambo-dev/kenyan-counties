package handler

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mambo-dev/kenya-locations/internal"
	"github.com/mambo-dev/kenya-locations/internal/database"
)

func transformToSubCountyResponse(subCounty database.SubCounty) internal.SubCountyResponse {

	subCountyID, err := uuid.Parse(subCounty.ID)

	if err != nil {
		// stops program as this should never happen
		log.Fatal("Failed to parse sub-county ID:", err)
		return internal.SubCountyResponse{}
	}

	return internal.SubCountyResponse{
		ID:               subCountyID,
		Name:             subCounty.Name,
		SubCountyGivenID: subCounty.SubCountyGivenID,
		CountyID:         subCounty.CountyID,
		CreatedAt:        subCounty.CreatedAt.Time,
		UpdatedAt:        subCounty.UpdatedAt.Time,
	}
}

func (h *Handler) GetSubCounties(w http.ResponseWriter, r *http.Request) {
	subCounties, err := h.cfg.Db.ListSubCounties(r.Context(), database.ListSubCountiesParams{
		Limit:  290,
		Offset: 0,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch sub-counties", err, false)
		return
	}

	response := make([]internal.SubCountyResponse, 0, len(subCounties))

	for _, subCounty := range subCounties {
		response = append(response, transformToSubCountyResponse(subCounty))
	}

	respondWithJSON(w, http.StatusOK, internal.ApiResponse{
		Status: "success",
		Data:   response,
	})
}

func (h *Handler) GetSubCountiesByCountyID(w http.ResponseWriter, r *http.Request) {

	countyIDStr := chi.URLParam(r, "countyID")
	if countyIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid county ID parameter", errors.New("invalid county ID passed"), false)
		return
	}
	countyID, err := uuid.Parse(countyIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid county ID format", err, false)
		return
	}

	subCounties, err := h.cfg.Db.GetSubCountiesByCountyID(r.Context(), database.GetSubCountiesByCountyIDParams{
		CountyID: countyID.String(),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "No sub-counties found for the specified county", err, false)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch sub-counties for the county", err, false)
		return
	}

	response := make([]internal.SubCountyResponse, 0, len(subCounties))

	for _, subCounty := range subCounties {
		response = append(response, transformToSubCountyResponse(subCounty))
	}

	respondWithJSON(w, http.StatusOK, internal.ApiResponse{
		Status: "success",
		Data:   response,
	})
}

func (h *Handler) GetSubCountyByName(w http.ResponseWriter, r *http.Request) {

	subCountyName := chi.URLParam(r, "subCountyName")
	if subCountyName == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid sub-county name parameter", errors.New("invalid sub-county name passed"), false)
		return
	}

	subCounty, err := h.cfg.Db.GetSubCountyByName(r.Context(), strings.ToLower(subCountyName))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Sub-county not found", err, false)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch sub-county by name", err, false)
		return
	}

	respondWithJSON(w, http.StatusOK, internal.ApiResponse{
		Status: "success",
		Data:   transformToSubCountyResponse(subCounty),
	})
}

func (h *Handler) SearchSubCountyByName(w http.ResponseWriter, r *http.Request) {

	subCountyName := r.URL.Query().Get("subCountyName")
	if subCountyName == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid sub-county name parameter", errors.New("invalid sub-county name passed"), false)
		return
	}
	if strings.Contains(subCountyName, `"`) || strings.Contains(subCountyName, "'") {
		respondWithError(w, http.StatusBadRequest, "Invalid sub-county name parameter", errors.New("invalid characters in sub-county name"), false)
		return
	}
	subCounties, err := h.cfg.Db.SearchSubCountiesByName(r.Context(), database.SearchSubCountiesByNameParams{
		LOWER:  subCountyName,
		Limit:  20,
		Offset: 0,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to search sub-counties by name", err, false)
		return
	}
	response := make([]internal.SubCountyResponse, 0, len(subCounties))
	for _, subCounty := range subCounties {
		response = append(response, transformToSubCountyResponse(subCounty))
	}

	respondWithJSON(w, http.StatusOK, internal.ApiResponse{
		Status: "success",
		Data:   response,
	})

}
