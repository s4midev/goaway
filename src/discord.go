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

	fmt.Println(cfg.Webhook)
	fmt.Println("GOT HERE!!!")
	client := webhook.NewClient(cfg.Webhook)

	embed := webhook.NewEmbed().SetTitle(title).SetDescription(body)

	fmt.Println(title)
	fmt.Println(body)

	err := client.SendEmbed(embed)

	if err != nil {
		fmt.Println(err)
	}
}
