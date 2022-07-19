package main

import (
	"log"
	"time"

	"github.com/iamgoangle/gokub/slack"
)

func main() {
	cfg := &slack.Configs{
		WebhookURL:    "https://hooks.slack.com/services/xxx",
		Timeout:       5 * time.Second,
		AppName:       "golf-test",
		RetryMax:      2,
		NextRetryTime: 5 * time.Second,
	}

	slackClient, err := slack.New(cfg)
	if err != nil {
		panic(err)
	}

	payload := &slack.Payload{
		Text: "Hello, world",
	}

	err = slackClient.SendMessage(payload)
	if err != nil {
		log.Println(err)
	}
}
