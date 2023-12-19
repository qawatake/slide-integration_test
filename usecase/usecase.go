package usecase

import (
	"context"
)

type Usecase struct {
	greetingClient Greeter
}

type Greeter interface {
	Hello(ctx context.Context) error
}

func New(greetingClient Greeter) *Usecase {
	return &Usecase{greetingClient: greetingClient}
}

func (u *Usecase) Do(ctx context.Context) error {
	return u.greetingClient.Hello(ctx)
}
