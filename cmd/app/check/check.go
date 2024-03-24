package check

import (
	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use: "check",
	Short: "Check bot stats",
	Long: "Check information about the currently authenticated bot",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	CheckCmd.AddCommand(CheckMeCmd)
	CheckCmd.AddCommand(CheckUpdatesCmd)
}

// vim: noexpandtab
