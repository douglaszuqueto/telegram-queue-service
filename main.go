package main

import (
	"log"
	"os"

	"github.com/douglaszuqueto/telegram/telegram"
)

func main() {
	config := telegram.Config{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		ChatID: os.Getenv("TELEGRAM_CHATID"),
	}

	bot := telegram.New(&config)

	_, err := bot.SendMessage("ok")

	if err != nil {
		log.Panic(err.Error())
	}
}
