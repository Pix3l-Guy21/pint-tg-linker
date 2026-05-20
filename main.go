package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// Check if the message contains the word "marco"
	// if not, return without doing anything
	if !strings.Contains(strings.ToLower(body.Message.Text), "marco") {
		return
	}

	if err := sayHello(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func sayHello(chatID int64) error {
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   "hello??",
	}
	reqByte, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	res, err := http.Post("https://api.telegram.org/bot7934825396:AAEW-hCX5aTTT7jhDS4S4iqiGq9eNcZCqds/sendMessage", "application/json", bytes.NewBuffer(reqByte))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status: " + res.Status)
	}
	return nil
}

func main() {
	http.ListenAndServe(":4590", http.HandlerFunc(Handler))
}
