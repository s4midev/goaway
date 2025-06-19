package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/devopsarr/radarr-go/radarr"
)

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
				if strings.Contains(m, "Caution") || strings.Contains(m, "unsupported extension") {
					fmt.Println("Release " + r.GetTitle() + " has a warning")
					client.QueueAPI.DeleteQueue(context.Background(), *r.Id).RemoveFromClient(true).Blocklist(true).Execute()
					notifyWebhook(cfg, "Radarr download `"+r.GetTitle()+"` was removed", m)
				}
			}
		}
	}
}
