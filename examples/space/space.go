package main

import (
	"fmt"
	"os"
	"path/filepath"

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

	file, _ := os.Create(filepath.Join("logo.png"))
	if err := api.GetSpaceIcon(file); err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	diskUsage, err := api.GetSpaceDiskUsage()
	if err != nil {
		fmt.Printf("on GetSpaceDiskUsage, error: %s\n", err)
		return
	}
	fmt.Printf("%#v \n", *diskUsage)
}
