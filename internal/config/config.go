package config

import (
	"dallyger/telegram-send/internal/telegram"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Location int

const (
	Any Location = iota
	Local
	User
	System
)

var config *Config

type Config struct {
	auth *viper.Viper
}

func InitConfig(aloc Location) (*Config, error) {

	a := viper.New()
	a.SetConfigName("auth")
	a.SetConfigType("toml")

	switch aloc {
	case Any:
		a.AddConfigPath(".")
		a.AddConfigPath("$HOME/.config/telegram-send")
		a.AddConfigPath("/etc/telegram-send")
	case Local:
		a.AddConfigPath(".")
	case User:
		a.AddConfigPath("$HOME/.config/telegram-send")
	case System:
		a.AddConfigPath("/etc/telegram-send")
	}

	if werr := a.SafeWriteConfig(); werr != nil {
		if _, ok := werr.(viper.ConfigFileAlreadyExistsError); ok {
			// this is okay. we wan to create a file and it does exist.
		} else {
			return nil, werr
		}
	}

	return GetConfig()
}

func GetConfig() (*Config, error) {

	if config != nil {
		return config, nil
	}

	auth, err := auth()
	if err != nil {
		return nil, err
	}

	config = &Config {
		auth: auth,
	}

	return config, nil
}

func (c Config) GetBot(botAlias string) (telegram.Bot, error) {
	if botAlias == "" {
		botAlias = "default"
	}

	bot := telegram.Bot {
		Id: c.auth.GetString(fmt.Sprintf("bot.%s", botAlias)),
	}

	if bot.Id == "" {
		return bot, errors.New(fmt.Sprintf("Bot [%s] is missing the auth token", botAlias))
	}

	return bot, nil
}

func (c Config) GetChat(chatAlias string) (telegram.Receiver, error) {
	if chatAlias == "" {
		chatAlias = "default"
	}

	id := telegram.Receiver(c.auth.GetString(fmt.Sprintf("chat.%s", chatAlias)))

	if id == "" {
		return id, errors.New(fmt.Sprintf("Chat [%s] is unknown", chatAlias))
	}

	return id, nil
}

func (c Config) SetBotAuth(token string, botAlias string) error {
	if botAlias == "" {
		botAlias = "default"
	}

	c.auth.Set(fmt.Sprintf("bot.%s", botAlias), token)
	return c.auth.WriteConfig()
}

func (c Config) SetChatAlias(id string, chatAlias string) error {
	if chatAlias == "" {
		chatAlias = "default"
	}

	c.auth.Set(fmt.Sprintf("chat.%s", chatAlias), id)
	return c.auth.WriteConfig()
}

func auth() (*viper.Viper, error) {

	v := viper.New()

	v.SetConfigName("auth")
	v.SetConfigType("toml")

	// check in the current working directory
	v.AddConfigPath(".")
	// check in the user's config directory
	v.AddConfigPath("$HOME/.config/telegram-send")
	// check in the system-wide config directory
	v.AddConfigPath("/etc/telegram-send")

	// load configuration from file
	err := v.ReadInConfig()

	return v, err
}

// vim: noexpandtab
