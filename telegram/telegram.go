package telegram

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Config config
type Config struct {
	Token  string
	ChatID string
	Debug  bool
}

// Telegram Telegram
type Telegram struct {
	bot    *tgbotapi.BotAPI
	chatID string
}

// New new
func New(config *Config) *Telegram {

	bot, err := tgbotapi.NewBotAPI(config.Token)

	if err != nil {
		log.Panic(err.Error())
	}

	bot.Debug = config.Debug

	return &Telegram{
		bot:    bot,
		chatID: config.ChatID,
	}
}

func (s *Telegram) chatIDToInt64() int64 {
	id, _ := strconv.Atoi(s.chatID)

	return int64(id)
}

// SendMessage SendMessage
func (s *Telegram) SendMessage(message string) (tgbotapi.Message, error) {

	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           s.chatIDToInt64(),
			ReplyToMessageID: 0,
		},
		ParseMode: "Markdown",
		Text:      message,
		DisableWebPagePreview: false,
	}

	return s.bot.Send(msg)
}

// Send Send
func (s *Telegram) Send(message tgbotapi.MessageConfig) (tgbotapi.Message, error) {
	return s.bot.Send(message)
}

// Subscribe subscribe
func (s *Telegram) Subscribe() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return s.bot.GetUpdatesChan(u)
}
