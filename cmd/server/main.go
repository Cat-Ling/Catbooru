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
	for name, clientCfg := range cfg.Clients {
		if clientCfg.Enabled {
			var client booru.BooruClient
			switch name {
			case "waifu_im":
				client = waifuim.New(clientCfg.BaseURL)
			case "waifu_pics":
				client = waifupics.New(clientCfg.BaseURL)
			case "nekos_moe":
				client = nekosmoe.New(clientCfg.BaseURL)
			case "nekos_api":
				client = nekosapi.New(clientCfg.BaseURL)
			case "pic_re":
				client = picre.New(clientCfg.BaseURL)
			case "purrbot_site":
				client = purrbotsite.New(clientCfg.BaseURL)
			default:
				log.Printf("Warning: Unknown client '%s'", name)
			}
			if client != nil {
				clients = append(clients, client)
			}
		}
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
