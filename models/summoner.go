package models

import "encoding/json"

var MarshalJSON = json.Marshal

type Summoner struct {
	AccountID     string `json:"accountId"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	Name          string `json:"name"`
	Id            string `json:"id"`
	Puuid         string `json:"puuid"`
	SummonerLevel int64  `json:"summonerLevel"`
}
