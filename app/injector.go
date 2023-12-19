//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"sample_app/usecase"

	"github.com/google/wire"
)

func New(
	context.Context,
) (*usecase.Usecase, error) {
	panic(wire.Build(
		CoreSet,
		configSet,
	))
}
