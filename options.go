package telegram_bot

import (
	"fmt"
	"os"
	"strconv"
)

type Options struct {
	ApiToken   string
	Debug      bool
	WebHookUrl string
	CertPath   string
	KeyPath    string
}

func FromEnv() Options {
	token := getEnvVar("BOT_TOKEN")
	webHookUrl := getEnvVar("BOT_WEB_HOOK_URL")
	certPath := getEnvVar("BOT_CERT_PATH")
	keyPath := getEnvVar("BOT_KEY_PATH")
	debugString := getEnvVar("BOT_DEBUG")
	debug := false
	debug, _ = strconv.ParseBool(debugString)

	return Options{
		ApiToken:   token,
		Debug:      debug,
		WebHookUrl: webHookUrl,
		CertPath:   certPath,
		KeyPath:    keyPath,
	}
}

func getEnvVar(varName string) string {
	value := os.Getenv(varName)
	if varName == "" {
		message, _ := fmt.Scanf("Env variable %s is not set", varName)
		panic(message)
	}
	return value
}
