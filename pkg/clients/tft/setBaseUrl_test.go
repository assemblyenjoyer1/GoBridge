package tft

import (
	"github.com/stretchr/testify/assert"
	"go-api-client/models"
	"testing"
)

func Test_SetBaseUrl(t *testing.T) {
	// and: test subject
	cfg := models.Config{
		BaseURL:       "sample-url",
		BaseRegionURL: "sample-region-url",
		ApiKey:        "test-key",
		Timeout:       10,
		Logger:        testData.Logger(t),
	}
	c := NewClient(cfg)

	// when: calling function
	c.SetBaseUrl("test-url")

	// then: no error returned
	assert.Equal(t, "test-url", c.baseURL)
}
