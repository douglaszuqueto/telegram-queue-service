package queue

import (
	"fmt"
	"log"

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
			Body: []byte(message),
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

// Stop Stop
func (s *Client) Stop() {
	log.Println("Closing rabbit channel")
	err := s.ch.Close()

	if err != nil {
		return
	}

	log.Println("Closing rabbit connection")

	err = s.conn.Close()

	if err != nil {
		return
	}
}
