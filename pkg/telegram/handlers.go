package telegram

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

const (
	commandStart = "start"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.bot.Send(msg)
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case commandStart:
		msg := b.createNewMessage(message, "Команда "+commandStart+" была введена верно!")
		b.bot.Send(msg)

	default:
		msg := b.createNewMessage(message, "Такой команды не существует")
		b.bot.Send(msg)
	}

}

func (b *Bot) createNewMessage(message *tgbotapi.Message, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(message.Chat.ID, text)
}
