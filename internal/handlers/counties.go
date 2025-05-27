package handler

import (
	"database/sql"
	"net/http"

	"github.com/mambo-dev/kenya-locations/config"
)

type Handler struct {
	db  *sql.DB
	cfg *config.APIConfig
}

func NewHandler(db *sql.DB, cfg *config.APIConfig) *Handler {
	return &Handler{db: db, cfg: cfg}
}

func (h *Handler) GetCounties(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) GetCounty(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) GetCountyByName(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) GetCountySubCounties(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, http.StatusNoContent, nil)
}
