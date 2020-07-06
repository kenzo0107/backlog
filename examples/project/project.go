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

	input := &backlog.GetProjectsOptions{}
	projects, _ := c.GetProjects(input)
	for _, project := range projects {
		fmt.Printf("project ID: %d, Name %s\n", *project.ID, *project.Name)
	}
}
