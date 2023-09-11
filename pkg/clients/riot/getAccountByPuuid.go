package riot

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

func (c *Client) GetAccountByPuuid(puuid string) (models.Account, *models.Response, error) {
	req := request.NewRequest(
		c.httpClient,
		http.MethodGet,
		fmt.Sprintf("%s/account/v1/accounts/by-puuid/%s", c.baseURL, puuid),
		c.apikey,
		bytes.NewBuffer(nil),
	)

	resp, err := req.Send()
	if err != nil {
		c.logger.With("response_code", resp.Code, "error", err).Error("get summoner by name failed")
		return models.Account{}, &models.Response{}, err
	}

	var account models.Account

	if resp.Code == http.StatusOK {
		if err = unmarshal(resp.Body, &account); err != nil {
			return models.Account{}, resp, err
		}
	}
	return account, resp, nil
}
