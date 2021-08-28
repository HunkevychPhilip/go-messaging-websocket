package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", h.Index).Methods(http.MethodGet, http.MethodPatch)
	router.HandleFunc("/ws-chat", h.Chat)

	return router
}
