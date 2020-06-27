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

	input := &backlog.GetUserMySelfRecentrlyViewedIssuesInput{
		Order:  backlog.OrderAsc,
		Offset: backlog.Int(0),
		Count:  backlog.Int(100),
	}
	issues, err := api.GetUserMySelfRecentrlyViewedIssues(input)

	if err != nil {
		fmt.Printf("(>_<) %s\n", err)
		return
	}

	for _, i := range issues {
		fmt.Printf("id: %d, issue key: %s, summary: %s\n", *i.Issue.ID, *i.Issue.IssueKey, *i.Issue.Summary)
	}
}
