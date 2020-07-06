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

	wiki, err := c.GetWiki(333028)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("wiki ID: %d, Name: %s", *wiki.ID, *wiki.Name)
}
