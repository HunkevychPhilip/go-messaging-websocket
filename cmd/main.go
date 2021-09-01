package main

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/datastore"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/handler"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/service"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/utils"
	"github.com/PhilipHunkevych/go-messaging-app/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatal(err.Error())
	}

	rw := utils.NewResponseWriter()
	rdb, err := datastore.NewRedisClient(
		viper.GetString("redis.host"),
		viper.GetString("redis.port"),
	)
	if err != nil {
		logrus.Fatal("Failed to create redis client" + err.Error())
	}
	db := datastore.NewDatastore(rdb)
	s := service.NewService(db)
	u := utils.NewUtils(rw)
	h := handler.NewHandler(u, s)

	port := viper.GetString("app.port")

	httpServ := server.NewHTTPServer(port, h.InitRoutes())

	logrus.Info("Attempting to start server...")
	logrus.Fatal(httpServ.Start())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
