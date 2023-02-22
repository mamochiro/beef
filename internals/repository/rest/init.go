package rest

import (
	"github.com/mamochiro/beef/internals/config"
	"github.com/mamochiro/beef/internals/infrastructure/httpClient"
)

type Client struct {
	httpClient *httpClient.HttpClient
	config     config.Configuration
}

func NewClient(httpClient *httpClient.HttpClient, config config.Configuration) Interface {
	return &Client{
		httpClient: httpClient,
		config:     config,
	}
}
