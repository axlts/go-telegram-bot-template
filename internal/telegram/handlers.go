package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	ErrUnknownCmd = errors.New("unknown command")
)

const (
	cmdStart = "start"
	cmdHelp  = "help"
)

// startText is the welcome message sent when the bot starts.
const startText = `
Welcome to your Telegram Bot!

Type /help for a list of available commands.
`

// helpText provides usage instructions for commands.
const helpText = `
Usage:
  /start    Initialize your session.
  /help     Display this help message.
`

// handleCmd processes incoming command messages.
func (bot *Bot) handleCmd(msg *tgbotapi.Message) {
	switch msg.Command() {
	case cmdStart:
		bot.send(msg.Chat.ID, startText)
	case cmdHelp:
		bot.send(msg.Chat.ID, helpText)
	default:
		bot.send(msg.Chat.ID, ErrUnknownCmd.Error())
	}
}

// handleMsg processes regular text messages from users.
func (bot *Bot) handleMsg(msg *tgbotapi.Message) {
	bot.send(msg.Chat.ID, msg.Text) // echo.
}
