package tft

import (
	"encoding/json"
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
	testStrings = fmt.Sprintf(`["random-match1", "random-match2", "random-match3"]`)
)

func Test_GetMatchListByPuuid_NoError(t *testing.T) {
	// given: test http client
	testClient := test.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(testStrings)),
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
	_, resp, err := c.GetMatchListByPuuid("test-puuid")

	// then: no error returned
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func Test_GetMatchListByPuuid_Unmarshal_Error(t *testing.T) {
	// given: test http client
	testClient := test.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(testStrings)),
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
	_, _, err := c.GetMatchListByPuuid("test-puuid")
	unmarshal = json.Unmarshal

	// then: no error returned
	assert.Error(t, err)
}

func Test_GetMatchListByPuuid_Send_Error(t *testing.T) {
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
	_, _, err := c.GetMatchListByPuuid("test-puuid")

	// then: error returned
	assert.Error(t, err)
}
