package main

import (
	"context"
	"log"
	"sample_app/app"
)

func main() {
	ctx := context.Background()
	u, err := app.New(ctx)
	if err != nil {
		log.Panicln(err)
	}
	if err := u.Do(context.Background()); err != nil {
		log.Panicln(err)
	}
}
