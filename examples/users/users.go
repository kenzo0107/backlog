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

	user, err := c.GetUserMySelf()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("user ID: %d, Name %s, userID: %s\n", *user.ID, *user.Name, *user.UserID)
}
