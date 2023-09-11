package tft

import (
	"go-api-client/models"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	apikey     string
	httpClient *http.Client
	logger     *zap.SugaredLogger
}

func NewClient(config models.Config) *Client {
	return &Client{
		baseURL: string(config.BaseRegionURL + "/tft"),
		apikey:  config.ApiKey,
		httpClient: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Minute,
		},
		logger: config.Logger,
	}
}
