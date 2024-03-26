package main

import (
	"dallyger/telegram-send/cmd/app/auth"
	"dallyger/telegram-send/cmd/app/chat"
	"dallyger/telegram-send/cmd/app/check"
	"dallyger/telegram-send/internal/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	animations []string
	audios     []string
	photos     []string
	videos     []string
	voices     []string

	files    []string
	messages []string

	rootCmd = &cobra.Command{
		Use:  "tg -m <message>",
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

			for _, animation := range animations {
				if err := bot.SendAnimation(chat, animation); err != nil {
					fmt.Fprintln(os.Stderr, err)
					defer os.Exit(1)
				}
			}

			for _, audio := range audios {
				if err := bot.SendAudio(chat, audio); err != nil {
					fmt.Fprintln(os.Stderr, err)
					defer os.Exit(1)
				}
			}

			for _, photo := range photos {
				if err := bot.SendPhoto(chat, photo); err != nil {
					fmt.Fprintln(os.Stderr, err)
					defer os.Exit(1)
				}
			}

			for _, video := range videos {
				if err := bot.SendVideo(chat, video); err != nil {
					fmt.Fprintln(os.Stderr, err)
					defer os.Exit(1)
				}
			}

			for _, voice := range voices {
				if err := bot.SendVoice(chat, voice); err != nil {
					fmt.Fprintln(os.Stderr, err)
					defer os.Exit(1)
				}
			}

			for _, file := range files {
				if err := bot.SendDocument(chat, file); err != nil {
					fmt.Fprintln(os.Stderr, err)
					defer os.Exit(1)
				}
			}

			for _, message := range messages {
				bot.SendMessage(chat, message)
			}

			return nil
		},
	}
)

func init() {
	rootCmd.Flags().StringArrayVar(&animations, "animation", nil, "send animation")
	rootCmd.Flags().StringArrayVarP(&audios, "audio", "a", nil, "send audio file")
	rootCmd.Flags().StringArrayVarP(&audios, "photo", "p", nil, "send image")
	rootCmd.Flags().StringArrayVar(&videos, "video", nil, "send video")
	rootCmd.Flags().StringArrayVar(&voices, "voice", nil, "send .ogg audio file as a voice message")

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
