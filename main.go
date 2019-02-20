package main

import (
	"log"
	"os"

	queueService "github.com/douglaszuqueto/telegram/queue"
	telegramService "github.com/douglaszuqueto/telegram/telegram"
)

func main() {
	configQueue := queueService.Config{
		IP:       os.Getenv("RABBITMQ_IP"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		Username: os.Getenv("RABBITMQ_USERNAME"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	}

	configTelegram := telegramService.Config{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		ChatID: os.Getenv("TELEGRAM_CHATID"),
	}

	telegram, err := telegramService.New(&configTelegram)
	if err != nil {
		log.Panic(err.Error())
	}

	queue, err := queueService.New(&configQueue)
	if err != nil {
		log.Panic(err.Error())
	}

	defer queue.CloseChannel()
	defer queue.Disconnect()

	err = queue.SendMessage("hellllo")
	if err != nil {
		log.Panic(err.Error())
	}

	messages, err := queue.Consume()
	if err != nil {
		log.Panic(err.Error())
	}

	for d := range messages {
		log.Printf("Received a message: %s", d.Body)
		telegram.SendMessage(string(d.Body))
		d.Ack(true)
	}

}
