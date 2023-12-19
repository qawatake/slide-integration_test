package app

import (
	"sample_app/adapter/greeting"
	"sample_app/config"
	"sample_app/usecase"

	"github.com/google/wire"
)

var CoreSet = wire.NewSet(
	usecase.New,
	wire.Bind(new(usecase.Greeter), new(*greeting.Client)),
	greeting.Set,
)

var configSet = wire.NewSet(
	config.NewGreeting,
)
