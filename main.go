package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var url, username, iconUrl, channel string

	flag.StringVar(&url, "url", "", "url to send webhooks")
	flag.StringVar(&username, "username", "", "custom name")
	flag.StringVar(&iconUrl, "icon-url", "", "custom icon")
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
		Icon:     iconUrl,
		Channel:  channel,
	}

	err := client.Send(w)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
