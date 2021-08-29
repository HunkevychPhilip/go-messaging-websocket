package handler

import (
	"github.com/spf13/viper"
	"net/http"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, viper.GetString("app.indexFilePath"))
}
