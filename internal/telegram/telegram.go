package telegram

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Bot struct {
	Id string
}

type Receiver string

func (b Bot) MeRaw () (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getMe", b.Id))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func (b Bot) GetUpdatesRaw () (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", b.Id))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func (b Bot) SendMessage(r Receiver, msg string) {
	resp, err := http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Id),
		url.Values{
			"chat_id": {string(r)},
			"text": {msg},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
}

// vim: noexpandtab
