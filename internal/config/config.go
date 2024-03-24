package config

import (
	"dallyger/telegram-send/internal/telegram"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	auth *viper.Viper
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

	// TODO: make internal once no component relies on it anymore.

	v := viper.New()

	v.SetConfigName("auth.toml")
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
