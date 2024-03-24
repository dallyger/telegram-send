package chat

import (
	"dallyger/telegram-send/internal/config"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

type Chat struct {
	Id    string
	Username string
	Label string
}

var ChatSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set default chat",
	Long:  "Retrieve and persist the default chat by checking for users who wrote with the bot.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.GetConfig()
		if err != nil {
			return err
		}

		bot, err := cfg.GetBot("default")
		if err != nil {
			return err
		}

		updates, err := bot.GetUpdates()
		if err != nil {
			return err
		}

		var choices = make(map[string]Chat)
		for _, update := range updates.Result {
			if update.Message != nil {
				chatId := fmt.Sprintf("%d", update.Message.From.Id)
				choices[chatId] = Chat{
					Id:    chatId,
					Username: update.Message.From.Username,
					Label: update.Message.From.Username,
				}
			}
		}

		if len(choices) == 0 {
			cmd.SilenceUsage = true
			return errors.New(`Please message the bot on Telegram, so that the chat can be auto-detected.`)
		}

		if len(choices) > 1 {
			// TODO: implement asking user to interactively choose default chat
			cmd.SilenceUsage = true
			return errors.New(`Too many users found`)
		}

		var selected Chat
		for _, chat := range choices {
			selected = chat
			break
		}

		if err := cfg.SetChatAlias(selected.Id, "default"); err != nil {
			return err
		}

		if err := cfg.SetChatAlias(selected.Id, selected.Username); err != nil {
			return err
		}

		return nil
	},
}

// vim: noexpandtab
