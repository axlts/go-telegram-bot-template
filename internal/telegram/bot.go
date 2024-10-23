package telegram

import (
	"context"
	"sync"

	"github.com/axlts/go-telegram-bot-template/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot represents the Telegram bot.
type Bot struct {
	api *tgbotapi.BotAPI
}

// New initializes a new Bot instance with the provided configuration.
func New(cfg config.Bot) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}
	api.Debug = cfg.Debug
	return &Bot{api: api}, nil
}

// Run starts the bot's main event loop, processing incoming updates from Telegram.
func (bot *Bot) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.api.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case update := <-updates:
			// ignore all non-message updates.
			if update.Message == nil {
				continue
			}
			// ignore all non-text messages.
			if update.Message.Text == "" {
				continue
			}

			if update.Message.IsCommand() {
				bot.handleCmd(update.Message)
			} else {
				bot.handleMsg(update.Message)
			}
		}
	}
}

// send sends a message to a specific chat.
func (bot *Bot) send(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := bot.api.Send(msg); err != nil {
		panic(err)
	}
}
