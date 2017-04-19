package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

const queue_name = "stats_processing"

type Config struct {
	Host string
	User string
	Pass string
	Port int
}

type Client struct {
	Config Config
	conn   *amqp.Connection
}

func (c *Client) Publish(data []byte) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(queue_name, true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         data,
		})
	return err
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// Constructs a new AMQP connection
func New(config Config) (*Client, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.User,
		config.Pass,
		config.Host,
		config.Port))
	if err != nil {
		return nil, err
	}
	c := &Client{
		Config: config,
		conn:   conn,
	}
	return c, nil
}
