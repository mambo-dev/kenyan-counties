package handler

import "net/http"

func (h *Handler) GetWards(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) GetWard(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNoContent, nil)
}
