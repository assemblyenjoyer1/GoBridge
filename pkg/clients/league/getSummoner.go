package league

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

func (c *Client) getSummoner(url string) (models.Summoner, *models.Response, error) {
	req := request.NewRequest(
		c.httpClient,
		http.MethodGet,
		fmt.Sprintf("%s/summoner/v4/summoners%s", c.baseURL, url),
		c.apikey,
		bytes.NewBuffer(nil),
	)

	resp, err := req.Send()
	if err != nil {
		c.logger.With("response_code", resp.Code, "error", err).Error("get summoner by name failed")
		return models.Summoner{}, &models.Response{}, err
	}

	var summoner models.Summoner

	if resp.Code == http.StatusOK {
		if err = unmarshal(resp.Body, &summoner); err != nil {
			return models.Summoner{}, resp, nil
		}
	}
	return summoner, resp, nil
}

func (c *Client) GetSummonerByEncryptedAccountId(encryptedAccountId string) (models.Summoner, *models.Response, error) {
	return c.getSummoner(fmt.Sprintf("/by-account/%s", encryptedAccountId))
}

func (c *Client) GetSummonerByPuuid(puuid string) (models.Summoner, *models.Response, error) {
	return c.getSummoner(fmt.Sprintf("/by-puuid/%s", puuid))
}

func (c *Client) GetSummonerByName(name string) (models.Summoner, *models.Response, error) {
	return c.getSummoner(fmt.Sprintf("/by-name/%s", name))
}
