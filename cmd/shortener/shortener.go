package main

import (
	"log"

	"github.com/junnotantra/go-shortener/internal/api"
	"github.com/junnotantra/go-shortener/internal/config"
)

func main() {
	var (
		err error
		cfg *config.Config
	)
	// initialize config
	err = config.Init()
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	cfg = config.Get()

	api.Main(cfg)
}
