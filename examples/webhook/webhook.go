package main

import (
	"fmt"
	"os"

	"github.com/kenzo0107/backlog"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")
	c := backlog.New(apiKey, baseURL)

	webhooks, err := c.GetWebhooks("SRE")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, webhook := range webhooks {
		fmt.Printf("id: %d, name: %s, hook url :%s\n", *webhook.ID, *webhook.Name, *webhook.HookURL)
	}
}
