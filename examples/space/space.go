package main

import (
	"fmt"
	"os"

	"github.com/kenzo0107/backlog"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")
	api := backlog.New(apiKey, baseURL, backlog.OptionDebug(true))

	space, err := api.GetSpace()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("space key: %s, name %s\n", space.SpaceKey, space.Name)
}
