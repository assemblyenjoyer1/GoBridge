package league

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

var testSummoner = fmt.Sprintf(`
											{
   												 "AccountID": "1",
   												 "ProfileIconId": 1,
   												 "RevisionDate": 1,
   												 "Name": "Test-name",
   												 "Id": "11",
   												 "Puuid": "111",
   												 "SummonerLevel": 300
											}
						`)

func Test_GetSummonerByName_Happy(t *testing.T) {
	// given: test http client
	testClient := test.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(testSummoner)),
			Header:     make(http.Header),
		}
	})

	// and: test subject
	cfg := models.Config{
		BaseURL: "sample-url",
		ApiKey:  "test-key",
		Timeout: 10,
		Logger:  testData.Logger(t),
	}

	c := NewClient(cfg)
	c.httpClient = testClient

	// when: calling function
	_, resp, err := c.GetSummonerByName("test-name")

	// then: no error returned
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func Test_GetSummonerByName_Unmarshal_Error(t *testing.T) {
	// given: test http client
	testClient := test.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(testSummoner)),
			Header:     make(http.Header),
		}
	})

	// and: test subject
	cfg := models.Config{
		BaseURL: "sample-url",
		ApiKey:  "test-key",
		Timeout: 10,
		Logger:  testData.Logger(t),
	}

	c := NewClient(cfg)
	c.httpClient = testClient

	unmarshal = func(data []byte, v any) error {
		return testError
	}

	// when: calling function
	_, resp, err := c.GetSummonerByEncryptedAccountId("test-name")
	unmarshal = json.Unmarshal

	// then: no error returned
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func Test_GetSummonerByName_Send_Error(t *testing.T) {
	// given: test http client
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}))
	defer srv.Close()
	testClient := srv.Client()

	// and: test subject
	cfg := models.Config{
		BaseURL: "sample-url",
		ApiKey:  "test-key",
		Timeout: 10,
		Logger:  testData.Logger(t),
	}

	c := NewClient(cfg)
	c.httpClient = testClient

	// when: calling function
	_, _, err := c.GetSummonerByPuuid("test-name")

	// then: error returned
	assert.Error(t, err)
}
