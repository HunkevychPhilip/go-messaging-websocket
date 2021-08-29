package main

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/handler"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/service"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/types"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/utils"
	"github.com/PhilipHunkevych/go-messaging-app/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	chatChannels := types.ChatChannels{
		ClientRequests:    make(chan *types.Client, 100),
		ClientDisconnects: make(chan string, 100),
		Messages:          make(chan *types.Msg, 100),
	}

	r := utils.NewResponseWriter()
	s := service.NewService(&chatChannels)
	u := utils.NewUtils(r)
	h := handler.NewHandler(u, s)


	port := viper.GetString("app.port")

	chatServ := server.NewChatServer(&chatChannels)
	httpServ := server.NewHTTPServer(port, h.InitRoutes())

	go chatServ.Start()
	log.Fatal(httpServ.Start())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
