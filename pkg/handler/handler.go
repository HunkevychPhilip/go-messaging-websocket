package handler

import "github.com/PhilipHunkevych/go-messaging-app/pkg/utils"

type Handler struct {
	utils *utils.Utils
}

func NewHandler(u *utils.Utils) *Handler {
	return &Handler{utils: u}
}
