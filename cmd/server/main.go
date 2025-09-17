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
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	var clients []booru.BooruClient
	if clientCfg, ok := cfg.Clients["waifu_im"]; ok && clientCfg.Enabled {
		clients = append(clients, waifuim.New(clientCfg.BaseURL))
	}
	if clientCfg, ok := cfg.Clients["waifu_pics"]; ok && clientCfg.Enabled {
		clients = append(clients, waifupics.New(clientCfg.BaseURL))
	}
	if clientCfg, ok := cfg.Clients["nekos_moe"]; ok && clientCfg.Enabled {
		clients = append(clients, nekosmoe.New(clientCfg.BaseURL))
	}
	if clientCfg, ok := cfg.Clients["nekos_api"]; ok && clientCfg.Enabled {
		// The docs say v4, but the config has v2. We'll use v4.
		clients = append(clients, nekosapi.New("https://api.nekosapi.com/v4"))
	}
	if clientCfg, ok := cfg.Clients["pic_re"]; ok && clientCfg.Enabled {
		clients = append(clients, picre.New(clientCfg.BaseURL))
	}
	if clientCfg, ok := cfg.Clients["purrbot_site"]; ok && clientCfg.Enabled {
		// The docs say v2, but the config has no version. We'll add it.
		clients = append(clients, purrbotsite.New(clientCfg.BaseURL+"/v2"))
	}

	if len(clients) == 0 {
		log.Fatal("No clients are enabled in the configuration.")
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	srv := server.New(clients)

	fmt.Printf("Starting server on %s with %d clients\n", addr, len(clients))
	if err := srv.Start(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
