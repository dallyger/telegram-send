package main

import (
	"dallyger/telegram-send/internal/config"
	"dallyger/telegram-send/internal/telegram"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {

	auth, err := config.Auth()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	bot, err := getBot(auth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	var messages []string

	cmd := &cobra.Command {
		Use: "tg user [user, ...]",
		Aliases: []string{"telegram-send"},
		Example: "tg @botfather /start",
		Long: "Send Telegram messages as a bot",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			errs := sendMessages(bot, auth, args, messages)
			if len(errs) > 0 {
				for _, err := range errs {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				}
				defer os.Exit(1)
			}
		},
	}

	cmd.Flags().StringArrayVarP(&messages, "msg", "m", nil, "message to send to user(s)")

	cmd.Execute()
}

func getBot(auth *viper.Viper) (telegram.Bot, error) {
	bot := telegram.Bot {
		Id: auth.GetString("bot.default"),
	}

	if bot.Id == "" {
		return bot, errors.New("Bot [default] is missing the auth token")
	}

	return bot, nil
}

func getChat(auth *viper.Viper, chat string) (telegram.Receiver, error) {
	id := telegram.Receiver(auth.GetString(fmt.Sprintf("chat.%s", chat)))

	if id == "" {
		return id, errors.New(fmt.Sprintf("Chat [%s] is unknown", chat))
	}

	return id, nil
}

func sendMessages(
	bot telegram.Bot,
	auth *viper.Viper,
	chats []string,
	messages []string,
) []error {
	errs := []error{}

	for _, chat := range chats {
		receiver, err := getChat(auth, chat)
		if err != nil {
			errs = append(errs, err)
		} else {
			for _, message := range messages {
				bot.SendMessage(receiver, message)
			}
		}
	}

	return errs
}

// vim: noexpandtab
