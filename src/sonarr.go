package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/devopsarr/sonarr-go/sonarr"
)

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
					notifyWebhook(cfg, "Sonarr download `"+r.GetTitle()+"` was removed", m)
				}
			}
		}
	}
}
