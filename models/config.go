package models

import (
	"go.uber.org/zap"
)

type Config struct {
	BaseURL       Region
	BaseRegionURL MatchRegion
	ApiKey        string
	Timeout       int64
	Logger        *zap.SugaredLogger
}
