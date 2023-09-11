package tft

import (
	"bytes"
	"fmt"
	"go-api-client/internal/request"
	"go-api-client/models"
	"net/http"
)

func (c *Client) GetMatchListByPuuid(puuid string) ([]string, *models.Response, error) {
	req := request.NewRequest(
		c.httpClient,
		http.MethodGet,
		fmt.Sprintf("%s/match/v1/matches/by-puuid/%s/ids", c.baseURL, puuid),
		c.apikey,
		bytes.NewBuffer(nil),
	)

	resp, err := req.Send()
	if err != nil {
		c.logger.With("response_code", resp.Code, "error", err).Error("get summoner by name failed")
		return nil, &models.Response{}, err
	}

	var matchIds []string

	if resp.Code == http.StatusOK {
		if err = unmarshal(resp.Body, &matchIds); err != nil {
			return matchIds, resp, err
		}
	}

	return matchIds, resp, nil
}
