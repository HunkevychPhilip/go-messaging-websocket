package datastore

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/types"
	"github.com/go-redis/redis"
)

type ChatRedis struct {
	rdb    *redis.Client
	pubsub *redis.PubSub
}

func NewChatRedis(rdb *redis.Client) *ChatRedis {
	return &ChatRedis{
		rdb: rdb,
	}
}

func (c *ChatRedis) Subscribe(channels ...string) (*redis.PubSub, error) {
	p := c.rdb.PSubscribe(channels...)

	_, err := p.Receive()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (c *ChatRedis) PostToChannel(channel string, msg types.Message) error {
	res := c.rdb.Publish(channel, msg)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}
