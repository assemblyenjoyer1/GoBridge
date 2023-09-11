package request

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errorReader struct{}

func (e errorReader) Read([]byte) (int, error) {
	return 0, errors.New("reader error")
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func Test_Send_Happy(t *testing.T) {
	// given: test http client
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`ok`)),
			Header:     make(http.Header),
		}
	})

	// and: test subject
	request := NewRequest(client, http.MethodGet, "test-url", "test-tey", nil)

	// when: sending request
	response, err := request.Send()

	// then: no error returned
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
}

func Test_Send_Request_Error(t *testing.T) {
	// given: test http client
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`ok`)),
			Header:     make(http.Header),
		}
	})

	// and: mocked http request function
	tmp := newRequest
	newRequest = func(method, url string, body io.Reader) (*http.Request, error) {
		return nil, errors.New("http request error")
	}

	// and: test subject
	request := NewRequest(client, http.MethodGet, "test-url", "test-key", nil)

	// when: sending request
	response, err := request.Send()

	// then: error returned
	assert.Error(t, err)
	assert.Nil(t, response.Body)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	newRequest = tmp
}

func Test_Send_Do_Error(t *testing.T) {
	// given: test http client
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}))
	defer server.Close()

	// and: test subject
	request := NewRequest(server.Client(), http.MethodGet, "test-url", "test-key", nil)

	// when: sending request
	response, err := request.Send()

	// then: error returned
	assert.Error(t, err)
	assert.Nil(t, response.Body)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

}

func Test_Send_Reader_Error(t *testing.T) {
	// given: test http client
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(errorReader{}),
			Header:     make(http.Header),
		}
	})

	// and: test subject
	request := NewRequest(client, http.MethodGet, "test-url", "test-token", nil)

	// when: sending request
	response, err := request.Send()

	// then: error returned
	assert.Error(t, err)
	assert.Nil(t, response.Body)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

}
