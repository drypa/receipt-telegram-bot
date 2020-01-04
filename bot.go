package telegram_bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
)

func Start(options Options) error {
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

	for update := range updatesChan {
		log.Printf("%v\n", update)
	}
	return nil
}
