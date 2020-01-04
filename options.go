package telegram_bot

import (
	"errors"
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

func (options Options) validate() error {
	err := validateEmpty(options.ApiToken, "Api token is not set")
	if err != nil {
		return err
	}
	err = validateEmpty(options.WebHookUrl, "Web hook url is not set")
	if err != nil {
		return err
	}
	err = validateEmpty(options.CertPath, "Certificate path is not set")
	if err != nil {
		return err
	}
	err = validateEmpty(options.KeyPath, "SSL key file path is not set")
	if err != nil {
		return err
	}
	return nil
}

func validateEmpty(value string, emptyErrorMessage string) error {
	if value == "" {
		return errors.New(emptyErrorMessage)
	}
	return nil
}

func getEnvVar(varName string) string {
	value := os.Getenv(varName)
	if varName == "" {
		message, _ := fmt.Scanf("Env variable %s is not set", varName)
		panic(message)
	}
	return value
}
