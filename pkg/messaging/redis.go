package messaging

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func NewRedisClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
