package tft

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-api-client/models"
	"go-api-client/test"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	testData  = test.NewTestData()
	testError = errors.New("")
)

var testMatch = fmt.Sprintf(`
{
    "metadata": {
        "dataVersion": "random-dataVersion",
        "matchId": "random-matchId",
        "participants": ["participant1", "participant2"]
    },
    "info": {
        "gameDateTime": 1234567890,
        "gameLength": 3600,
        "gameVariation": "random-variation",
        "gameVersion": "random-version",
        "participants": [
            {
                "companion": {
                    "skinId": 1,
                    "contentId": "random-contentId",
                    "species": "random-species"
                },
                "goldLeft": 50,
                "lastRound": 10,
                "level": 8,
                "placement": 4,
                "playersEliminated": 2,
                "puuid": "random-puuid",
                "timeEliminated": 1234567890,
                "totalDamageToPlayers": 500,
                "traits": [
                    {
                        "name": "Trait1",
                        "numUnits": 2,
                        "style": 1,
                        "tierCurrent": 3,
                        "tierTotal": 4
                    },
                    {
                        "name": "Trait2",
                        "numUnits": 3,
                        "style": 2,
                        "tierCurrent": 2,
                        "tierTotal": 5
                    }
                ],
                "units": [
                    {
                        "items": ["item1", "item2"],
                        "characterId": "character1",
                        "chosen": "random-chosen",
                        "name": "Unit1",
                        "rarity": 3,
                        "unit": 2
                    },
                    {
                        "items": ["item3", "item4"],
                        "characterId": "character2",
                        "chosen": "random-chosen",
                        "name": "Unit2",
                        "rarity": 2,
                        "unit": 1
                    }
                ]
            }
        ],
        "queueId": 420,
        "tftSetNumber": 5
    }
}
`)

func TestClient_GetMatchByMatchId_NoError(t *testing.T) {
	// given: test http client
	testClient := test.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(testMatch)),
			Header:     make(http.Header),
		}
	})

	// and: test subject
	cfg := models.Config{
		BaseURL:       "sample-url",
		BaseRegionURL: "random-region-url",
		ApiKey:        "test-key",
		Timeout:       10,
		Logger:        testData.Logger(t),
	}

	c := NewClient(cfg)
	c.httpClient = testClient

	// when: calling function
	_, resp, err := c.GetMatchByMatchId("test-name")

	// then: no error returned
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestClient_GetMatchByMatchId_Unmarshal_Error(t *testing.T) {
	// given: test http client
	testClient := test.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(testMatch)),
			Header:     make(http.Header),
		}
	})

	// and: test subject
	cfg := models.Config{
		BaseURL:       "sample-url",
		BaseRegionURL: "random-region-url",
		ApiKey:        "test-key",
		Timeout:       10,
		Logger:        testData.Logger(t),
	}

	unmarshal = func(data []byte, v any) error {
		return testError
	}

	c := NewClient(cfg)
	c.httpClient = testClient

	// when: calling function
	_, _, err := c.GetMatchByMatchId("test-name")
	unmarshal = json.Unmarshal

	// then: no error returned
	assert.Error(t, err)
}

func Test_GetMatchByMatchId_Send_Error(t *testing.T) {
	// given: test http client
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}))
	defer srv.Close()
	testClient := srv.Client()

	// and: test subject
	cfg := models.Config{
		BaseURL:       "sample-url",
		BaseRegionURL: "sample-region-url",
		ApiKey:        "test-key",
		Timeout:       10,
		Logger:        testData.Logger(t),
	}

	c := NewClient(cfg)
	c.httpClient = testClient

	// when: calling function
	_, _, err := c.GetMatchByMatchId("test-name")

	// then: error returned
	assert.Error(t, err)
}
