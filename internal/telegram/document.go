package telegram

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (b Bot) SendAnimation(chat Receiver, filePath string) error {
	return b.makeDocumentRequest("sendAnimation", chat, filePath, "animation")
}

func (b Bot) SendAudio(chat Receiver, filePath string) error {
	return b.makeDocumentRequest("sendAudio", chat, filePath, "audio")
}

func (b Bot) SendDocument(chat Receiver, filePath string) error {
	return b.makeDocumentRequest("sendDocument", chat, filePath, "document")
}

func (b Bot) SendPhoto(chat Receiver, filePath string) error {
	return b.makeDocumentRequest("sendPhoto", chat, filePath, "photo")
}

func (b Bot) SendVideo(chat Receiver, filePath string) error {
	return b.makeDocumentRequest("sendVideo", chat, filePath, "video")
}

func (b Bot) SendVoice(chat Receiver, filePath string) error {
	return b.makeDocumentRequest("sendVoice", chat, filePath, "voice")
}

func (b Bot) makeDocumentRequest(tgAction string, chat Receiver, filePath string, fileProp string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)

	// text part1
	p1w, _ := bw.CreateFormField("chat_id")
	p1w.Write([]byte(string(chat)))

	// file part1
	_, fileName := filepath.Split(filePath)
	fw1, _ := bw.CreateFormFile(fileProp, fileName)
	io.Copy(fw1, file)

	bw.Close()
	file.Close()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.Id, tgAction), buf)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", bw.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}

	return nil
}

// vim: noexpandtab
