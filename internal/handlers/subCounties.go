package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mambo-dev/kenya-locations/internal/database"
)

func (h *Handler) GetSubCounties(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
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

	_, err = h.cfg.Db.GetSubCountiesByCountyID(r.Context(), database.GetSubCountiesByCountyIDParams{
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

	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) GetSubCountyByName(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) SearchSubCountyByName(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}
