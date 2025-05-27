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
	"github.com/mambo-dev/kenya-locations/internal/utils"
)

func transformToWardResponse(ward database.Ward) internal.WardResponse {
	wardID, err := uuid.Parse(ward.ID)

	if err != nil {
		// stops program as this should never happen
		log.Fatal("Failed to parse ward ID:", err)
		return internal.WardResponse{}
	}

	return internal.WardResponse{
		ID:          wardID,
		Name:        ward.Name,
		WardGivenID: ward.WardGivenID,
		SubCountyID: ward.SubCountyID,
		CreatedAt:   ward.CreatedAt.Time,
		UpdatedAt:   ward.UpdatedAt.Time,
	}
}

func (h *Handler) GetWards(w http.ResponseWriter, r *http.Request) {
	wardsCount, err := h.cfg.Db.TotalWards(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to count wards", err, false)
		return
	}
	limitOffset, errMsg, err := utils.SafeLimitOffsetParser(r, wardsCount)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, errMsg, err, false)
		return
	}
	wards, err := h.cfg.Db.ListWards(r.Context(), database.ListWardsParams{
		Limit:  int64(limitOffset.Limit),
		Offset: int64(limitOffset.Offset),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch wards", err, false)
		return
	}
	wardResponses := make([]internal.WardResponse, 0, len(wards))
	for _, ward := range wards {
		wardResponses = append(wardResponses, transformToWardResponse(ward))
	}
	respondWithJSON(w, http.StatusOK, internal.ApiResponse{
		Status: "success",
		Data: internal.PaginatedResponse{
			TotalCount: wardsCount,
			Data:       wardResponses,
		},
	})
}

func (h *Handler) GetWardByName(w http.ResponseWriter, r *http.Request) {
	wardName := chi.URLParam(r, "wardName")
	if wardName == "" {
		respondWithError(w, http.StatusBadRequest, "Ward name is required", nil, false)
		return
	}

	ward, err := h.cfg.Db.GetWardByName(r.Context(), wardName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Ward not found", nil, false)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch ward", err, false)
		return
	}
	wardResponse := transformToWardResponse(ward)
	respondWithJSON(w, http.StatusOK, wardResponse)
}

func (h *Handler) SearchWardByName(w http.ResponseWriter, r *http.Request) {
	wardName := r.URL.Query().Get("wardName")
	if wardName == "" {
		respondWithError(w, http.StatusBadRequest, "Ward name is required for search", nil, false)
		return
	}

	if strings.Contains(wardName, `"`) || strings.Contains(wardName, "'") {
		respondWithError(w, http.StatusBadRequest, "Invalid ward name parameter", errors.New("invalid characters in ward name"), false)
		return
	}

	wardsCount, err := h.cfg.Db.TotalWards(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to count wards", err, false)
		return
	}

	limitOffset, errMsg, err := utils.SafeLimitOffsetParser(r, wardsCount)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, errMsg, err, false)
		return
	}

	wards, err := h.cfg.Db.SearchWardsByName(r.Context(), database.SearchWardsByNameParams{
		LOWER:  wardName,
		Limit:  int64(limitOffset.Limit),
		Offset: int64(limitOffset.Offset),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to search wards", err, false)
		return
	}

	wardResponses := make([]internal.WardResponse, 0, len(wards))
	for _, ward := range wards {
		wardResponses = append(wardResponses, transformToWardResponse(ward))
	}

	respondWithJSON(w, http.StatusOK, internal.ApiResponse{
		Status: "success",
		Data: internal.PaginatedResponse{
			TotalCount: wardsCount,
			Data:       wardResponses,
		},
	})
}

func (h *Handler) GetWardsBySubCountyId(w http.ResponseWriter, r *http.Request) {
	subCountyIdStr := chi.URLParam(r, "subCountyID")
	if subCountyIdStr == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid sub county ID parameter", errors.New("invalid sub county ID passed"), false)
		return
	}

	countyID, err := uuid.Parse(subCountyIdStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid sub county ID format", err, false)
		return
	}

	wards, err := h.cfg.Db.GetWardsBySubCountyID(r.Context(), database.GetWardsBySubCountyIDParams{
		SubCountyID: countyID.String(),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "No wards found for the specified county", err, false)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch wards for the county", err, false)
		return
	}

	response := make([]internal.WardResponse, 0, len(wards))

	for _, ward := range wards {
		response = append(response, transformToWardResponse(ward))
	}

	respondWithJSON(w, http.StatusOK, internal.ApiResponse{
		Status: "success",
		Data:   response,
	})
}
