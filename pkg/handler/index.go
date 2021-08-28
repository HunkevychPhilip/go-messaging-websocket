package handler

import (
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(viper.GetString("app.indexFilePath"))
	if err != nil {
		h.utils.ResponseWriter.GenerateError(w, http.StatusInternalServerError, "Failed to read 'index.html'.")

		return
	}

	if _, err = io.Copy(w, f); err != nil {
		h.utils.ResponseWriter.GenerateError(w, http.StatusInternalServerError, "Failed to send 'index.html' content.")

		return
	}
}
