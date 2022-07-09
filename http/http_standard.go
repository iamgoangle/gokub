package http

import (
	"net/http"
	"time"
)

// StandardHTTPConfigs represent http standard configs type
type StandardHTTPConfigs struct {
	MaxIdleConns       int
	IdleConnTimeout    time.Duration
	DisableCompression bool

	TimeoutInSec int
}

// NewStandardHTTP new http
func NewStandardHTTP(cfg *StandardHTTPConfigs) *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       cfg.MaxIdleConns,
		IdleConnTimeout:    cfg.IdleConnTimeout * time.Second,
		DisableCompression: cfg.DisableCompression,
	}

	h := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(cfg.TimeoutInSec) * time.Second,
	}

	return h
}
