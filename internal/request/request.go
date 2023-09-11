package request

import (
	"fmt"
	"go-api-client/models"
	"io"
	"net/http"
)

var (
	newRequest = http.NewRequest
)

type Request struct {
	client *http.Client
	method string
	url    string
	apikey string
	data   io.Reader
}

func NewRequest(client *http.Client, method, url, apikey string, data io.Reader) *Request {
	return &Request{
		client: client,
		method: method,
		url:    url,
		apikey: apikey,
		data:   data,
	}
}

func (r *Request) Send() (*models.Response, error) {
	request, err := newRequest(r.method, r.url, r.data)
	if err != nil {
		return &models.Response{
			Body: nil,
			Code: http.StatusInternalServerError,
		}, err
	}

	request.Header.Set("X-Riot-Token", r.apikey)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	request.Close = true

	res, err := r.client.Do(request)
	if err != nil {
		return &models.Response{
			Body: nil,
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("error while sending http request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &models.Response{
			Body: nil,
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("error while reading http response: %v", err)
	}

	return &models.Response{
		Body: body,
		Code: res.StatusCode,
	}, nil

}
