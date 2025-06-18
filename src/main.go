package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/devopsarr/radarr-go/radarr"
	sonarr "github.com/devopsarr/sonarr-go/sonarr"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	RadarrUrl     string `envconfig:"RADARR_URL" required:"true"`
	RadarrApiKey  string `envconfig:"RADARR_API_KEY" required:"true"`
	SonarrUrl     string `envconfig:"SONARR_URL" required:"true"`
	SonarrApiKey  string `envconfig:"SONARR_API_KEY" required:"true"`
	CheckInterval int    `envconfig:"CHECK_INTERVAL" default:"5000"`
}

func scanRadarr(cfg *Config) {
	config := &radarr.Configuration{
		Servers:       radarr.ServerConfigurations{{URL: cfg.RadarrUrl}},
		DefaultHeader: map[string]string{"Authorization": cfg.RadarrApiKey},
	}

	client := radarr.NewAPIClient(config)

	result, _, err := client.QueueAPI.GetQueue(context.Background()).Status([]radarr.QueueStatus{radarr.QUEUESTATUS_COMPLETED}).Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range result.GetRecords() {
		for _, s := range r.StatusMessages {
			fmt.Println(s.GetTitle())
			for _, m := range s.GetMessages() {
				fmt.Println(m)
				if strings.Contains(m, "Caution") {
					fmt.Println("Release " + r.GetTitle() + " has a warning")
					client.QueueAPI.DeleteQueue(context.Background(), *r.Id).RemoveFromClient(true).Blocklist(true).Execute()
				}
			}
		}
	}
}

func scanSonarr(cfg *Config) {
	config := &sonarr.Configuration{
		Servers:       sonarr.ServerConfigurations{{URL: cfg.SonarrUrl}},
		DefaultHeader: map[string]string{"Authorization": cfg.SonarrApiKey},
	}

	client := sonarr.NewAPIClient(config)

	result, _, err := client.QueueAPI.GetQueue(context.Background()).Status([]sonarr.QueueStatus{sonarr.QUEUESTATUS_COMPLETED}).Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range result.GetRecords() {
		for _, s := range r.StatusMessages {
			fmt.Println(s.GetTitle())
			for _, m := range s.GetMessages() {
				fmt.Println(m)
				if strings.Contains(m, "Caution") || strings.Contains(m, "unsupported extension") {
					fmt.Println("Episode " + r.GetTitle() + " has a warning")
					client.QueueAPI.DeleteQueue(context.Background(), *r.Id).RemoveFromClient(true).Blocklist(true).Execute()
				}
			}
		}
	}
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
