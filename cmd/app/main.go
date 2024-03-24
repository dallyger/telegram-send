package main

import (
	"dallyger/telegram-send/internal/config"

	"github.com/spf13/cobra"
)

var (
	messages []string

	rootCmd = &cobra.Command {
		Use: "tg user [user, ...]",
		Aliases: []string{"telegram-send"},
		Example: "tg @botfather /start",
		Long: "Send Telegram messages as a bot",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.GetConfig()
			if err != nil {
				return err
			}

			bot, err := cfg.GetBot("default")
			if err != nil {
				return err
			}

			chat, err := cfg.GetChat("default")
			if err != nil {
				return err
			}

			for _, message := range messages {
				bot.SendMessage(chat, message)
			}

			return nil
		},
	}
)

func init() {
	rootCmd.Flags().StringArrayVarP(&messages, "msg", "m", nil, "message to send to user(s)")
}

func main() {
	rootCmd.Execute()
}

// vim: noexpandtab
