package handler

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mambo-dev/kenya-locations/config"
	"github.com/mambo-dev/kenya-locations/internal"
	"github.com/mambo-dev/kenya-locations/internal/database"
	"github.com/mambo-dev/kenya-locations/internal/utils"
)

type Handler struct {
	db  *sql.DB
	cfg *config.APIConfig
}

func NewHandler(db *sql.DB, cfg *config.APIConfig) *Handler {
	return &Handler{db: db, cfg: cfg}
}

func transformToCountyResponse(county database.County) internal.CountyResponse {

	countyID, err := uuid.Parse(county.ID)

	if err != nil {
		// stops program as this should never happen
		log.Fatal("Failed to parse county ID:", err)
		return internal.CountyResponse{}
	}

	return internal.CountyResponse{
		ID:       countyID,
		Name:     county.Name,
		CountyID: county.CountyGivenID,
	}

}

func (h *Handler) GetCounties(w http.ResponseWriter, r *http.Request) {

	countiesCount, err := h.cfg.Db.TotalCounties(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to count counties", err, false)
		return
	}

	limitOffset, errMsg, err := utils.SafeLimitOffsetParser(r, countiesCount)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, errMsg, err, false)
		return
	}

	counties, err := h.cfg.Db.ListCounties(r.Context(), database.ListCountiesParams{
		Limit:  int64(limitOffset.Limit),
		Offset: int64(limitOffset.Offset),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch counties", err, false)
		return
	}

	response := make([]internal.CountyResponse, 0, len(counties))

	for _, county := range counties {
		response = append(response, transformToCountyResponse(county))
	}

	respondWithJSON(w, http.StatusOK, internal.ApiResponse{
		Status: "success",
		Data:   response,
	})
}

func (h *Handler) GetCountyByName(w http.ResponseWriter, r *http.Request) {

	countyName := chi.URLParam(r, "countyName")

	if countyName == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid county name parameter", errors.New("invalid county name passed"), false)
		return
	}

	county, err := h.cfg.Db.GetCountyByName(r.Context(), countyName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "County not found", err, false)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch county by name", err, false)
		return
	}

	respondWithJSON(w, http.StatusOK, internal.ApiResponse{
		Status: "success",
		Data:   transformToCountyResponse(county),
	})
}

func (h *Handler) SearchCountyByName(w http.ResponseWriter, r *http.Request) {

	countiesCount, err := h.cfg.Db.TotalCounties(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to count counties", err, false)
		return
	}

	limitOffset, errMsg, err := utils.SafeLimitOffsetParser(r, countiesCount)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, errMsg, err, false)
		return
	}

	countyName := r.URL.Query().Get("countyName")
	if countyName == "" {
		respondWithError(w, http.StatusBadRequest, "Missing countyName query parameter", errors.New("missing countyName parameter"), false)
		return
	}

	if strings.Contains(countyName, `"`) || strings.Contains(countyName, "'") {
		respondWithError(w, http.StatusBadRequest, "Invalid county name parameter", errors.New("invalid characters in county name"), false)
		return
	}

	counties, err := h.cfg.Db.SearchCountiesByName(r.Context(), database.SearchCountiesByNameParams{
		LOWER:  strings.ToLower(countyName),
		Limit:  int64(limitOffset.Limit),
		Offset: int64(limitOffset.Offset),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "No counties found matching the name", err, false)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to search counties by name", err, false)
		return
	}

	response := make([]internal.CountyResponse, 0, len(counties))
	for _, county := range counties {
		response = append(response, transformToCountyResponse(county))
	}
	respondWithJSON(w, http.StatusOK, internal.ApiResponse{
		Status: "success",
		Data:   response,
	})
}
