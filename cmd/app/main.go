package main

import (
	"dallyger/telegram-send/cmd/app/auth"
	"dallyger/telegram-send/cmd/app/chat"
	"dallyger/telegram-send/cmd/app/check"
	"dallyger/telegram-send/internal/config"

	"github.com/spf13/cobra"
)

var (
	files []string
	messages []string

	rootCmd = &cobra.Command {
		Use: "tg -m <message>",
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


			for _, file := range files {
				bot.SendDocument(chat, file)
			}

			for _, message := range messages {
				bot.SendMessage(chat, message)
			}

			return nil
		},
	}
)

func init() {
	rootCmd.Flags().StringArrayVarP(&files, "file", "f", nil, "send file")
	rootCmd.Flags().StringArrayVarP(&messages, "msg", "m", nil, "send message")

	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(chat.ChatCmd)
	rootCmd.AddCommand(check.CheckCmd)
}

func main() {
	rootCmd.Execute()
}

// vim: noexpandtab
