package main

import (
	"booru-server/internal/clients/nekosapi"
	"booru-server/internal/clients/nekosmoe"
	"booru-server/internal/clients/picre"
	"booru-server/internal/clients/purrbotsite"
	"booru-server/internal/clients/waifuim"
	"booru-server/internal/clients/waifupics"
	"booru-server/internal/config"
	"booru-server/internal/server"
	"booru-server/pkg/booru"
	"fmt"
	"log"
)

// @title Go Booru Server API
// @version 1.0
// @description A Go server for a feature-rich booru browsing site.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	clients := []booru.BooruClient{
		waifuim.New(),
		waifupics.New(),
		nekosmoe.New(),
		nekosapi.New(),
		picre.New(),
		purrbotsite.New(),
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	srv := server.New(clients, cfg.Auth, cfg.RateLimit)

	fmt.Printf("Starting server on %s with %d clients\n", addr, len(clients))
	if err := srv.Start(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
