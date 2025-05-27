package handler

import "net/http"

func (h *Handler) GetWards(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) GetWardByName(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) SearchWardByName(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) GetWardsBySubCountyId(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}
