package telegram_bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
)

func Start(options Options) error {
	err := options.validate()
	if err != nil {
		log.Println("Bot options invalid")
		return err
	}

	bot, err := tgbotapi.NewBotAPI(options.ApiToken)
	if err != nil {
		log.Println("Bot create error")
		return err
	}
	bot.Debug = options.Debug

	log.Printf("Autorised as %s", bot.Self.UserName)
	config := tgbotapi.NewWebhookWithCert(options.WebHookUrl+bot.Token, options.CertPath)
	_, err = bot.SetWebhook(config)
	if err != nil {
		log.Println("Web hook create error")
		return err
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Println("Web hook error")
		return err
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed. %s\n", info.LastErrorMessage)
	}
	updatesChan := bot.ListenForWebhook("/" + bot.Token)

	go http.ListenAndServeTLS(":8443", options.CertPath, options.KeyPath, nil)

	processUpdates(updatesChan, bot)
	return nil
}

func processUpdates(updatesChan tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updatesChan {
		log.Printf("%v\n", update)
		if update.Message == nil {
			continue
		}
		var responseText string
		switch update.Message.Text {
		case "":
			responseText = "Please enter a command."
		case "/start":
			responseText = "I'm a bot to collect Your purchase tickets."
		case "/register":
			register(update.Message.From.ID)
			responseText = "You are registered. I collect only your virtual telegram Id."
		default:
			err := tryAddReceipt(update.Message.From.ID, update.Message.Text)
			responseText = "Added"
			if err != nil {
				responseText = err.Error()
			}
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
		bot.Send(msg)
	}
}

func tryAddReceipt(userId int, messageText string) error {
	//TODO: validate receipt query string and store
	panic("not implemented exception")
}

func register(userId int) {
	//TODO: store user to DB
}
