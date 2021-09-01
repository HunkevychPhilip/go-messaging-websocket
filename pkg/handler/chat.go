package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) Chat(rw http.ResponseWriter, r *http.Request) {
	conn, err := h.service.Chat.UpgradeConn(rw, r)
	if err != nil {
		h.utils.ResponseWriter.GenerateError(rw, http.StatusInternalServerError, err.Error())

		return
	}

	logrus.Info("New conn successfully established.")

	h.service.Chat.ServeNewConn(conn)
}
