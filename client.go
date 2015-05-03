package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	Url string
}

type Webhook struct {
	Text     string `json:"text"`
	Username string `json:"username,omitempty"`
	Icon     string `json:"icon_url,omitempty"`
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

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Something bad happened: %s", body)
	}

	fmt.Println("Message sent")

	return nil
}
