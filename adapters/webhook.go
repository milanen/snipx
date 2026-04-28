package adapters

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendWebhook(url string, content string) {
	payload := map[string]string{
		"content": content,
	}

	jsonValue, _ := json.Marshal(payload)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
}
