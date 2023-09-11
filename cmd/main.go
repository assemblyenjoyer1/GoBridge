package main

import (
	"fmt"
	client "go-api-client"
	"go-api-client/models"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	l := &zap.SugaredLogger{}

	cfg := models.Config{
		BaseURL:       models.EUW,
		BaseRegionURL: models.EUROPE,
		ApiKey:        "RGAPI-e018676b-a05c-4154-9a75-017decae095a",
		Timeout:       10,
		Logger:        l,
	}
	lc := client.NewClient(cfg)
	res, resp, err := lc.League.GetSummonerByName("Grandmaster TFT")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	if resp.Code != http.StatusOK {
		fmt.Printf("error with the http response %v", resp.Code)
	}

	fmt.Println(res.Puuid)

	account, resp, err := lc.Riot.GetAccountByPuuid(res.Puuid)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	if resp.Code != http.StatusOK {
		fmt.Printf("error with the http response %v\n", resp.Code)
	}

	fmt.Println(account.Puuid + " - " + account.GameName)

	matchIds, resp, err := lc.Tft.GetMatchListByPuuid(account.Puuid)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	if resp.Code != http.StatusOK {
		fmt.Printf("error with the http response %v\n", resp.Code)
	}

	fmt.Println(matchIds[1])

	match, resp, err := lc.Tft.GetMatchByMatchId(matchIds[1])
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	if resp.Code != http.StatusOK {
		fmt.Printf("error with the http response %v\n", resp.Code)
	}

	fmt.Println(match.Metadata.MatchId)
	for i, v := range match.Info.Participants {
		res2, _, _ := lc.League.GetSummonerByPuuid(v.Puuid)
		fmt.Printf("%d: "+res2.Name+"\n", i)
	}

}
