package test

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"net/http"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func NewTestData() *TestData {
	return &TestData{
		AccountID:     "1",
		ProfileIconId: 1,
		RevisionDate:  1,
		Name:          "test-name",
		Id:            "1",
		Puuid:         "1",
		SummonerLevel: 300,
	}
}

type TestData struct {
	AccountID     string
	ProfileIconId int
	RevisionDate  int64
	Name          string
	Id            string
	Puuid         string
	SummonerLevel int64
}

func (d *TestData) Logger(t zaptest.TestingT) *zap.SugaredLogger {
	return zaptest.NewLogger(t).Sugar()
}
