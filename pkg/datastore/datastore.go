package datastore

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/types"
	"github.com/go-redis/redis"
)

type Chat interface {
	Subscribe(...string) (*redis.PubSub, error)
	PostToChannel(channel string, msg types.Message) error
}

type Datastore struct {
	Chat
}

func NewDatastore(rdb *redis.Client) *Datastore {
	return &Datastore{
		Chat: NewChatRedis(rdb),
	}
}
