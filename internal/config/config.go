package config

import (
	"github.com/spf13/viper"
)

func Auth() (*viper.Viper, error) {
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
