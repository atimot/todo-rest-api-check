package main

import (
	"context"
	"log"

	"github.com/atimot/app/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	run(ctx, cfg)
}

func run(ctx context.Context, cfg *config.Config) {

}
