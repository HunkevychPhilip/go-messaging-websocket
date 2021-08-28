package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", h.index).Methods(http.MethodGet, http.MethodPatch)

	return router
}
