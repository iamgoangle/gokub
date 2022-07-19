package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

var (
	WebhookEndpoint = "https://hooks.slack.com"
)

type Configs struct {
	WebhookURL  string
	AppName     string
	Environment string

	Timeout       time.Duration
	RetryMax      int
	NextRetryTime time.Duration
}

type Slack struct {
	cfg *Configs
}

type Payload struct {
	Text string `json:"text,omitempty"`
}

func New(cfg *Configs) (*Slack, error) {
	if len(cfg.WebhookURL) == 0 {
		return nil, errors.New("[New]: invalid configs")
	}

	return &Slack{
		cfg: cfg,
	}, nil
}

func (s *Slack) SendMessage(payload *Payload) error {
	request := gorequest.New().Timeout(s.cfg.Timeout)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-app-name", s.cfg.AppName)

	resp, _, err := request.
		Post(s.cfg.WebhookURL).
		Send(payload).
		Retry(s.cfg.RetryMax, s.cfg.NextRetryTime, http.StatusInternalServerError).
		End()
	if err != nil {
		return fmt.Errorf("[Slack.SendMessage]: unable to send message %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[Slack.SendMessage]: got http status %d", resp.StatusCode)
	}

	return nil
}
