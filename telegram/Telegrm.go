package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Telegram(message string) {
	bot, err := tgbotapi.NewBotAPI("1185283164:AAGo7ZzqddWaOrn30jTkE3w-dLVbElmKD2w")
	if err != nil {
		panic(err)
	}
	bot.Debug = false
	bot.Send(tgbotapi.NewMessage(-740163933, message))

}
