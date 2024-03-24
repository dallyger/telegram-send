package chat

import (
	"github.com/spf13/cobra"
)

var ChatCmd = &cobra.Command{
    Use: "chat",
    Long: "Manage chats",
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

func init() {
    ChatCmd.AddCommand(ChatSetCmd)
}
