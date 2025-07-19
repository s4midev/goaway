package main

import (
	"fmt"

	"github.com/versai-pro/discord-go/webhook"
)

func notifyWebhook(cfg *Config, title string, body string) {
	if cfg.Webhook == "" {
		fmt.Println(cfg.Webhook)
		return
	}

	client := webhook.NewClient(cfg.Webhook)

	embed := webhook.NewEmbed().SetTitle(title).SetDescription(body)

	err := client.SendEmbed(embed)

	if err != nil {
		fmt.Println(err)
	}
}
