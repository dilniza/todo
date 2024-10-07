package pkg

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func TelegramBotFunc(msg interface{}) (string, error) {

	botToken := "7027904551:AAEgUYBeU9Aar0tVi5lEzLUwjK3LJ0HqfNA"
	chatID := "-4194845131"

	messageBytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	message := string(messageBytes)

	payload := struct {
		ChatID string `json:"chat_id"`
		Text   string `json:"text"`
	}{
		ChatID: chatID,
		Text:   message,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := "https://api.telegram.org/bot" + botToken + "/sendMessage"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	success := "Response Status:" + resp.Status
	return success, nil
}

func GenerateOTP() int {
	

	return rand.Intn(900000) + 100000
}
