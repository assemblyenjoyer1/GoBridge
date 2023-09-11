package client

import (
	"go-api-client/models"
	"go-api-client/pkg/clients/league"
	"go-api-client/pkg/clients/riot"
	"go-api-client/pkg/clients/tft"
)

type Client struct {
	League LeagueClient
	Riot   RiotClient
	Tft    TftClient
}

func NewClient(config models.Config) *Client {
	return &Client{
		League: league.NewClient(config),
		Riot:   riot.NewClient(config),
		Tft:    tft.NewClient(config),
	}
}
