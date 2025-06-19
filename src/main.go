package main

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	RadarrUrl     string `envconfig:"RADARR_URL" required:"true"`
	RadarrApiKey  string `envconfig:"RADARR_API_KEY" required:"true"`
	SonarrUrl     string `envconfig:"SONARR_URL" required:"true"`
	SonarrApiKey  string `envconfig:"SONARR_API_KEY" required:"true"`
	CheckInterval int    `envconfig:"CHECK_INTERVAL" default:"5000"`
}

func main() {
	fmt.Println("Go away! :3")

	var cfg Config
	envconfig.MustProcess("GOAWAY", &cfg)

	ticker := time.NewTicker(time.Duration(cfg.CheckInterval) * time.Millisecond)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			fmt.Println("Scanning Radarr and Sonarr :P")
			scanRadarr(&cfg)
			scanSonarr(&cfg)
		}
	}()

	select {}
}
