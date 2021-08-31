package main

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/handler"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/messaging"
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

	rw := utils.NewResponseWriter()
	r, err := messaging.NewRedisClient()
	if err != nil {
		log.Fatal("Failed to create redis client" + err.Error())
	}
	m := messaging.NewMessaging(r)
	s := service.NewService(m)
	u := utils.NewUtils(rw)
	h := handler.NewHandler(u, s)

	port := viper.GetString("app.port")

	httpServ := server.NewHTTPServer(port, h.InitRoutes())

	log.Fatal(httpServ.Start())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
