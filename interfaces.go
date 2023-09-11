package client

import "go-api-client/models"

type LeagueClient interface {
	GetSummonerByEncryptedAccountId(encryptedAccountId string) (models.Summoner, *models.Response, error)
	GetSummonerByPuuid(puuid string) (models.Summoner, *models.Response, error)
	GetSummonerByName(name string) (models.Summoner, *models.Response, error)
	SetBaseUrl(baseURL string)
}

type RiotClient interface {
	GetAccountByPuuid(puuid string) (models.Account, *models.Response, error)
	SetBaseUrl(baseURL string)
}

type TftClient interface {
	GetMatchByMatchId(matchId string) (models.Match, *models.Response, error)
	GetMatchListByPuuid(puuid string) ([]string, *models.Response, error)
	SetBaseUrl(baseURL string)
}
