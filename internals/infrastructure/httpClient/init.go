package httpClient

import (
	"github.com/go-resty/resty/v2"
	"github.com/mamochiro/beef/internals/config"
)

type HttpClient struct {
	Config config.Configuration
	Client *resty.Client
}

func NewHttpClient(config config.Configuration) *HttpClient {
	client := resty.New().SetBaseURL(config.BeefEndpoint)
	return &HttpClient{
		Config: config,
		Client: client,
	}
}
