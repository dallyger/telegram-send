package main

import (
	"bufio"
	"dallyger/telegram-send/cmd/app/auth"
	"dallyger/telegram-send/cmd/app/chat"
	"dallyger/telegram-send/cmd/app/check"
	"dallyger/telegram-send/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev-snapshot"

	animations []string
	audios     []string
	photos     []string
	videos     []string
	voices     []string

	files    []string
	messages []string
	stdin    bool

	rootCmd = &cobra.Command{
		Use:  "tg -m <message>",
		Long: "Send Telegram messages as a bot",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				cmd.Help()
				return nil
			}

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

			args = os.Args[1:]
			for {
				if len(args) == 0 {
					break
				}

				var err error
				switch args[0] {

				case "--animation":
					err = bot.SendAnimation(chat, args[1])
					args = args[2:]

				case "-a", "--audio":
					err = bot.SendAudio(chat, args[1])
					args = args[2:]

				case "-f", "--file":
					err = bot.SendDocument(chat, args[1])
					args = args[2:]

				case "-m", "--msg":
					bot.SendMessage(chat, args[1])
					args = args[2:]

				case "-p", "--photo":
					err = bot.SendPhoto(chat, args[1])
					args = args[2:]

				case "--stdin":
					err = bot.SendStream(chat, bufio.NewReader(os.Stdin))
					args = args[1:]

				case "--video":
					err = bot.SendVideo(chat, args[1])
					args = args[2:]

				case "--voice":
					err = bot.SendVoice(chat, args[1])
					args = args[2:]

				default:
					log.Fatalf("Invalid arg: %+v\n", args[0])
				}

				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					defer os.Exit(1)
				}
			}

			return nil
		},
	}
)

func init() {
	rootCmd.Flags().StringArrayVar(&animations, "animation", nil, "send animation")
	rootCmd.Flags().StringArrayVarP(&audios, "audio", "a", nil, "send audio file")
	rootCmd.Flags().StringArrayVarP(&photos, "photo", "p", nil, "send image")
	rootCmd.Flags().StringArrayVar(&videos, "video", nil, "send video")
	rootCmd.Flags().StringArrayVar(&voices, "voice", nil, "send .ogg audio file as a voice message")

	rootCmd.Flags().StringArrayVarP(&files, "file", "f", nil, "send file")
	rootCmd.Flags().StringArrayVarP(&messages, "msg", "m", nil, "send message")
	rootCmd.Flags().BoolVar(&stdin, "stdin", false, "send message by reading stdin")

	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(chat.ChatCmd)
	rootCmd.AddCommand(check.CheckCmd)
}

func main() {
	rootCmd.Execute()
}

// vim: noexpandtab
