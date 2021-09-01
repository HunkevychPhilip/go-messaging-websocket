package datastore

import (
	"fmt"
	"github.com/go-redis/redis"
)

func NewRedisClient(host, port string) (*redis.Client, error) {
	rdb := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%s", host, port),
		},
	)

	_, err := rdb.Ping().Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
