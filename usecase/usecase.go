package usecase

import (
	"context"
	"sample_app/adapter/greeting"
)

type Usecase struct {
	greetingClient *greeting.Client
}

func New(greetingClient *greeting.Client) *Usecase {
	return &Usecase{greetingClient: greetingClient}
}

func (u *Usecase) Greet(ctx context.Context) error {
	return u.greetingClient.Hello(ctx)
}
