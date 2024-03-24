package check

import (
	"dallyger/telegram-send/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

var CheckUpdatesCmd = &cobra.Command{
	Use: "updates",
	Short: "Check bot updates",
	Long: "Raw dump of the getUpdates API endpoint",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.GetConfig()
		if err != nil {
			return err
		}

		bot, err := cfg.GetBot("default")
		if err != nil {
			return err
		}

		msg, err := bot.GetUpdatesRaw()
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
