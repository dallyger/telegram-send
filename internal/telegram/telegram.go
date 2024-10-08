package telegram

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
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

// Send messages in chunks if it is too long.
func (b Bot) SendMessageChunked(r Receiver, msg string) {
	charLimit := 4096

	for {
		idx := -1

		if len(msg) <= charLimit {
			// full message fits in a single request
			idx = len(msg)
		}

		if idx == -1 {
			// try to split messages on newlines
			idx = strings.LastIndex(msg[:charLimit], "\n")
		}

		if idx == -1 {
			// no newline in chunk. send as much as possible.
			idx = charLimit
		}

		b.SendMessage(r, msg[:idx])
		msg = msg[idx:]

		if msg == "" {
			// continue until nothing is left to send
			break
		}
	}
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

func (b Bot) SendStream(r Receiver, s *bufio.Reader) (error) {
	buf := ""

	for {
		line, err := s.ReadString('\n')
		buf += line

		if len(buf) > 1024*1024*8 {
			// buffer grows too big (8 mb); send and flush it.
			b.SendMessageChunked(r, buf)
			buf = ""
		}

		if err == nil {
			// continue reading
			continue
		}

		if buf != "" {
			// we've read it all. time to send.
			b.SendMessageChunked(r, buf)
			buf = ""
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
	}
	
	return nil
}

// vim: noexpandtab
