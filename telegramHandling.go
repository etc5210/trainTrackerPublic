package main

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SEND_MESSAGE(apiToken string, chatID int64, timeMessage string) (sendStatus int) {
	bot, err := tgbotapi.NewBotAPI(apiToken)
    if err != nil {
        log.Panic(err)
    }
    // Create a message to send
    message := tgbotapi.NewMessage(chatID, timeMessage)

    // Send the message
    _, err = bot.Send(message)
    if err != nil {
        log.Panic(err)
        sendStatus = 0
        return sendStatus
    } else {
        sendStatus = 1
        return sendStatus
    }
}
