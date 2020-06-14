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

	input := &backlog.GetIssuesInput{
		ProjectIDs: []int{82387, 82437, 31759},
		Count:      100,
	}
	issues, err := api.GetIssues(input)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, issue := range issues {
		fmt.Printf("id: %d, key: %s, summary: %s, CustomFields: %#v\n", issue.ID, issue.IssueKey, issue.Summary, issue.CustomFields)
	}
}
