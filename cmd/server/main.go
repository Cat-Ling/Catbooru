package main

import (
	"booru-server/internal/config"
	"booru-server/internal/modules/nekosapi"
	"booru-server/internal/modules/nekosmoe"
	"booru-server/internal/modules/picre"
	"booru-server/internal/modules/purrbotsite"
	"booru-server/internal/modules/waifuim"
	"booru-server/internal/modules/waifupics"
	"booru-server/internal/server"
	"booru-server/pkg/booru"
	"fmt"
	"log"
)

var version = "dev"

// @title Go Booru Server API
// @version 1.0
// @description A Go server for a feature-rich booru browsing site.
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	modules := []booru.BooruModule{
		waifuim.New(),
		waifupics.New(),
		nekosmoe.New(),
		nekosapi.New(),
		picre.New(),
		purrbotsite.New(),
	}

	if len(modules) == 0 {
		log.Fatal("No modules are enabled.")
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	srv := server.New(modules, version)

	fmt.Printf("Starting server version %s on %s with %d modules\n", version, addr, len(modules))
	if err := srv.Start(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
