package models

type Unit struct {
	Items       []string `json:"items"`
	CharacterId string   `json:"characterId"`
	Chosen      string   `json:"chosen"`
	Name        string   `json:"name"`
	Rarity      int      `json:"rarity"`
	Tier        int      `json:"unit"`
}

type Trait struct {
	Name        string `json:"name"`
	NumUnits    int    `json:"NumUnits"`
	Style       int    `json:"style"`
	TierCurrent int    `json:"tierCurrent"`
	TierTotal   int    `json:"tierTotal"`
}

type Companion struct {
	SkinId    int    `json:"skinId"`
	ContentId string `json:"contentId"`
	Species   string `json:"species"`
}

type Participant struct {
	Companion            Companion
	GoldLeft             int     `json:"goldLeft"`
	LastRound            int     `json:"lastRound"`
	Level                int     `json:"level"`
	Placement            int     `json:"placement"`
	PlayersEliminated    int     `json:"playersEliminated"`
	Puuid                string  `json:"puuid"`
	TimeEliminated       int64   `json:"timeEliminated"`
	TotalDamageToPlayers int     `json:"totalDamageToPlayers"`
	Traits               []Trait `json:"traits"`
	Units                []Unit  `json:"units"`
}

type Info struct {
	GameDateTime  int32         `json:"gameDateTime"`
	GameLength    int64         `json:"gameLength"`
	GameVariation string        `json:"gameVariation"`
	GameVersion   string        `json:"gameVersion"`
	Participants  []Participant `json:"participants"`
	QueueId       int           `json:"queueId"`
	TftSetNumber  int64         `json:"tftSetNumber"`
}

type Metadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type Match struct {
	Metadata Metadata
	Info     Info
}
