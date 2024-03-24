package auth

import (
	"dallyger/telegram-send/internal/config"
	"errors"
	"fmt"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var ask bool
var token string

var AuthCmd = &cobra.Command{
    Use: "auth",
    Long: "Authenticate as a bot",
    Short: "Authenticate as a bot",
    RunE: func(cmd *cobra.Command, args []string) error {
        config, err := config.GetConfig()
        if err != nil {
            return err
        }

        if token == "" && ask {
            var err error
            fmt.Print("Enter token [hidden]: ")
            token, err = readToken()
            if err != nil {
                return err
            }
        }

        if token == "" {
            return errors.New("Token missing")
        }

        return config.SetBotAuth(token, "default")
    },
}

func init() {
    AuthCmd.Flags().BoolVarP(&ask, "ask", "a", false, "prompt for token")
    AuthCmd.Flags().StringVarP(&token, "token", "t", "", "token used to authenticate the bot")
}

func readToken() (string, error) {
    bytes, err := term.ReadPassword(syscall.Stdin)
    if err != nil {
        return "", errors.New(fmt.Sprintf("Failed to read token: %s", err))
    }
    return strings.TrimSpace(string(bytes)), nil
}
