package main

import (
	"fmt"
	"os"

	"github.com/kenzo0107/backlog"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")
	api := backlog.New(apiKey, baseURL)

	user, err := api.GetUserMySelf()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("user ID: %d, Name %s\n", user.ID, user.Name)
}
