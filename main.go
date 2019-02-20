package main

import (
	"log"
	"os"

	telegramService "github.com/douglaszuqueto/telegram/telegram"
)

func main() {
	config := telegramService.Config{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		ChatID: os.Getenv("TELEGRAM_CHATID"),
	}

	telegram := telegramService.New(&config)

	_, err := telegram.SendMessage("ok")

	if err != nil {
		log.Panic(err.Error())
	}

}
