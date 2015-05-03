package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	Url string
}

type Webhook struct {
	Text     string `json:"text"`
	Username string `json:"username,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

type Client struct {
	Config Config
}

func New(url string) *Client {
	return &Client{
		Config: Config{
			Url: url,
		},
	}
}

func (self *Client) Send(webhook Webhook) error {
	j, err := json.Marshal(webhook)
	if err != nil {
		return err
	}

	resp, err := http.Post(self.Config.Url, "application/json", bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
