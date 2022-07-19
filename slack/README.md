# Slack

Slack is a Golang library to send messages to Slack Webhook

## Installation

```
go get github.com/iamgoangle/gokub/slack
```

## Usage

```go
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
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)