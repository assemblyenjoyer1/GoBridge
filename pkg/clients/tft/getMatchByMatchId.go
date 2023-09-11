package tft

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-client/internal/request"
	"go-api-client/models"
	"net/http"
)

var (
	unmarshal = json.Unmarshal
)

func (c *Client) GetMatchByMatchId(matchId string) (models.Match, *models.Response, error) {
	req := request.NewRequest(
		c.httpClient,
		http.MethodGet,
		fmt.Sprintf("%s/match/v1/matches/%s", c.baseURL, matchId),
		c.apikey,
		bytes.NewBuffer(nil),
	)

	resp, err := req.Send()
	if err != nil {
		c.logger.With("response_code", resp.Code, "error", err).Error("get summoner by name failed")
		return models.Match{}, &models.Response{}, err
	}

	var match models.Match

	if resp.Code == http.StatusOK {
		if err = unmarshal(resp.Body, &match); err != nil {
			return models.Match{}, resp, err
		}
	}

	return match, resp, nil
}
