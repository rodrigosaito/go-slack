package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var url, username, channel string

	flag.StringVar(&url, "url", "", "url to send webhooks")
	flag.StringVar(&username, "username", "", "custom name")
	flag.StringVar(&channel, "channel", "", "#channel or @user to send the message to")
	flag.Parse()

	message := flag.Arg(0)
	if message == "" {
		fmt.Println("You must specify a message")
		os.Exit(1)
	}

	if url == "" {
		fmt.Println("-url flag must be set")
	}

	client := New(url)

	w := Webhook{
		Text:     message,
		Username: username,
		Channel:  channel,
	}

	client.Send(w)
}
