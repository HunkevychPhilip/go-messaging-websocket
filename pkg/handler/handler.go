package handler

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/service"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/utils"
)

type Handler struct {
	service *service.Service
	utils   *utils.Utils
}

func NewHandler(u *utils.Utils, s *service.Service) *Handler {
	return &Handler{
		utils:   u,
		service: s,
	}
}
