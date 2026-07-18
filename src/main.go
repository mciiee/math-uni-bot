package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	godotenv "github.com/joho/godotenv"
)

const telegramTokenEnvVar = "TELEGRAM_BOT_TOKEN"

func getToken() string {
	err := godotenv.Load()
	
	if err != nil {
    log.Print("Failed to load .env, falling back to checking the environment variables")
	}
  
	token, present := os.LookupEnv(telegramTokenEnvVar)

	if !present {
		log.Printf("\"%s\" is not present in the environment", telegramTokenEnvVar)
    os.Exit(1)
  }

	return token
}

func main() {
	token := getToken()

  bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
    log.Panic("Failed to init the bot: ", err)
	}

  bot.Debug = true

	log.Printf("Authorized on account \"%s\"", bot.Self.UserName)

	updateConf := tgbotapi.NewUpdate(0)
	updateConf.Timeout = 60
  
	updates := bot.GetUpdatesChan(updateConf)

	for update := range updates {

    if update.Message == nil {
      continue
		}
    
		log.Printf("[Message:%d] %s<%d>: \"%s\"\n", update.Message.Chat.ID, update.Message.From.UserName,  update.Message.From.ID, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "\u2A01")
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
