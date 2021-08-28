package main

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/handler"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/utils"
	"github.com/PhilipHunkevych/go-messaging-app/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	rw := utils.NewResponseWriter()
	u := utils.NewUtils(rw)
	h := handler.NewHandler(u)
	s := server.NewServer(viper.GetString("app.port"), h.InitRoutes())

	log.Fatal(s.Start())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
