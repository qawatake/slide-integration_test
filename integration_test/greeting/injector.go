//go:build wireinject
// +build wireinject

package greeting

import (
	"context"
	"sample_app/app"
	"sample_app/config"
	"sample_app/usecase"

	"github.com/google/wire"
)

func New(
	context.Context,
	config.Greeting,
) (*usecase.Usecase, error) {
	panic(wire.Build(
		app.CoreSet,
	))
}
