package cafa

import (
	"net"
	"net/http"
	"time"
)

type Client struct {
	baseUrl    string
	httpClient *http.Client
}

func NewClient(baseURL string, customHttpCli *http.Client) *Client {

	const (
		MaxIdleConns    = 1
		MaxConnsPerHost = 1
		timeout         = 5 * time.Second
	)
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,

		ExpectContinueTimeout: 4 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,

		MaxIdleConns:    MaxIdleConns,
		MaxConnsPerHost: MaxConnsPerHost,
	}

	cli := Client{
		baseUrl: baseURL,
		httpClient: &http.Client{
			Transport: tr,
			Timeout:   timeout,
		},
	}
	if customHttpCli != nil {
		cli.httpClient = customHttpCli
	}
	return &cli
}
