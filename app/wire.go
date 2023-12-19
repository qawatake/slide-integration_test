package app

import (
	"sample_app/adapter/greeting"
	"sample_app/config"
	"sample_app/usecase"

	"github.com/google/wire"
)

var CoreSet = wire.NewSet(
	usecase.New,
	greeting.New,
	greeting.NewHTTPClient,
)

var configSet = wire.NewSet(
	config.NewGreetingConfig,
)
