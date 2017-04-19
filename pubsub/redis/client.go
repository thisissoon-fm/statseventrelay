package redis

import (
	"context"
	"sync"

	"statseventrelay/pubsub"

	redis "gopkg.in/redis.v5"
)

// Redis Pub/Sub configuration
type Config struct {
	Host string
}

// Redis Pub/Sub client
type Client struct {
	Config Config
	wg     sync.WaitGroup
	redis  *redis.Client // Underlying redis client
}

// Subscribe to a topic, returning a channel that received messages
// will be placed on or an error
func (c *Client) Subscribe(ctx context.Context, topics ...string) (<-chan pubsub.Message, error) {
	sub, err := c.redis.PSubscribe(topics...)
	if err != nil {
		return nil, err
	}
	ch := make(chan pubsub.Message)
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		defer close(ch)
		defer sub.Close()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := sub.ReceiveMessage()
				if err != nil {
					return
				}
				ch <- pubsub.Message{
					Topic:   msg.Channel,
					Payload: []byte(msg.Payload),
				}
			}
		}
	}()
	return (<-chan pubsub.Message)(ch), nil
}

// Close redis client
func (c *Client) Close() error {
	defer c.wg.Wait()
	return c.redis.Close()
}

// Construct a new Redis Pub/Sub client
func New(config Config) *Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.Host,
	})
	return &Client{
		Config: config,
		redis:  client,
	}
}
