package check

import (
	"dallyger/telegram-send/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

var CheckMeCmd = &cobra.Command{
	Use: "me",
	Short: "Check bot stats",
	Long: "Check information about the currently authenticated bot",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.GetConfig()
		if err != nil {
			return err
		}

		bot, err := cfg.GetBot("default")
		if err != nil {
			return err
		}

		msg, err := bot.MeRaw()
		if err != nil {
			return err
		}

		fmt.Println(msg)
		return nil
	},
}

func init() {
}

// vim: noexpandtab
