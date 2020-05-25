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

	project, err := api.GetProject("SRE")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("project ID: %d, Name %s\n", project.ID, project.Name)
}
