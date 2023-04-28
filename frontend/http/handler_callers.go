package http

import (
	"github.com/aldarisbm/ltmllm/message"
	"net/http"
)

func Message() string {
	msg := message.Message{}
	client := &http.Client{}

	resp, err := client.Get("https://google.com")
}
