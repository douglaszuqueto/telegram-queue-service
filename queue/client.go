package queue

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Config config
type Config struct {
	Username string
	Password string
	IP       string
	Port     string
}

// Client client
type Client struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func makeURL(config *Config) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s", config.Username, config.Password, config.IP, config.Port)
}

// New new
func New(config *Config) (*Client, error) {
	url := makeURL(config)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn: conn,
		ch:   ch,
	}

	return client, nil
}

// SendMessage SendMessage
func (s *Client) SendMessage(message string) error {
	q, err := s.ch.QueueDeclare(
		"telegram", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)

	if err != nil {
		return err
	}

	err = s.ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	return nil
}

// Consume Consume
func (s *Client) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := s.ch.Consume(
		"telegram", // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		return nil, err
	}

	return msgs, nil
}

// Disconnect Disconnect
func (s *Client) Disconnect() {
	s.conn.Close()
}

// CloseChannel CloseChannel
func (s *Client) CloseChannel() {
	s.ch.Close()
}
