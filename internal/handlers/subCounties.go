package handler

import "net/http"

func (h *Handler) GetSubCounty(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}
func (h *Handler) GetSubCounties(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) GetSubCountyByName(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) GetSubCountyWards(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}
