package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Chat struct {
	Id         int    `json:"id"`
	Type       string `json:"type"`
	Title      string `json:"title"`
	Username   string `json:"username"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

type User struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsPremium    bool   `json:"is_premium"`
}

type Message struct {
	MessageId             int    `json:"message_id"`
	MessageThreadId       int    `json:"message_thread_id"`
	From                  *User  `json:"from"`
	Date                  int    `json:"date"`
	Chat                  *Chat  `json:"chat"`
	NewChatTitle          string `json:"new_chat_title"`
	GroupChatCreated      bool   `json:"group_chat_created"`
	SupergroupChatCreated bool   `json:"supergroup_chat_created"`
	ChannelChatCreated    bool   `json:"channel_chat_created"`
}

type Update struct {
	UpdateId          int      `json:"update_id"`
	Message           *Message `json:"message"`
	EditedMessage     *Message `json:"edited_message"`
	ChannelPost       *Message `json:"channel_post"`
	EditedChannelPost *Message `json:"edited_channel_post"`
}

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

func (b Bot) GetUpdates() (UpdatesResponse, error) {

	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", b.Id))
	if err != nil {
		return UpdatesResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var u UpdatesResponse

	if err := json.Unmarshal(body, &u); err != nil {
		return u, err
	}

	return u, nil

}

// vim: noexpandtab
