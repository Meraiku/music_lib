package main

import (
	"context"
	"log"

	"github.com/meraiku/music_lib/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("error creating new app: %s", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("unexpected server shutdown: %s", err)
	}
}
