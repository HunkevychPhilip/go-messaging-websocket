package messaging

import (
	"github.com/go-redis/redis"
)

type Messaging struct {
	*redis.Client
}

func NewMessaging(client *redis.Client) *Messaging {
	return &Messaging{
		Client: client,
	}
}
