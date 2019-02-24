package main

import (
	"fmt"
	"log"
	"os"
	"time"

	queueService "github.com/douglaszuqueto/telegram/queue"
)

func main() {
	configQueue := queueService.Config{
		IP:       os.Getenv("RABBITMQ_IP"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		Username: os.Getenv("RABBITMQ_USERNAME"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	}

	queue, err := queueService.New(&configQueue)
	if err != nil {
		log.Panic(err.Error())
	}

	defer queue.Stop()

	var counter int

	for {
		msg := fmt.Sprintf("*Message*: %v", counter)

		err = queue.SendMessage(msg)
		if err != nil {
			log.Panic(err.Error())
		}

		log.Printf("Sending message: %v", msg)
		counter++

		time.Sleep(5 * time.Second)
	}
}
