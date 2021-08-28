package main

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/handler"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/service"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/utils"
	"github.com/PhilipHunkevych/go-messaging-app/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	r := utils.NewResponseWriter()
	s := service.NewService()
	u := utils.NewUtils(r)
	h := handler.NewHandler(u, s)

	go s.ChatService.Router()
	serv := server.NewHTTPServer(viper.GetString("app.port"), h.InitRoutes())

	log.Fatal(serv.Start())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
