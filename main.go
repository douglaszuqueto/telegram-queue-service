package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	queueService "github.com/douglaszuqueto/telegram/queue"
	telegramService "github.com/douglaszuqueto/telegram/telegram"
)

func main() {
	done := make(chan bool)
	signalCh := make(chan os.Signal, 1)

	signal.Notify(signalCh, os.Interrupt)
	signal.Notify(signalCh, syscall.SIGTERM)

	go func() {
		<-signalCh
		log.Println("Stopping services")
		done <- true
	}()

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

	go func() {
		messages, err := queue.Consume()
		if err != nil {
			log.Panic(err.Error())
		}

		for {
			select {
			case message := <-messages:
				msg := string(message.Body)

				log.Println(msg)

				_, err := telegram.SendMessage(msg)
				if err != nil {
					log.Println(err.Error())
				}
				message.Ack(true)
			case <-time.After(10 * time.Second):
				log.Println("No messages")
			}
		}
	}()

	<-done
	queue.Stop()
	log.Println("Exited...")
}
