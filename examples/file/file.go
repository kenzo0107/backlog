package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kenzo0107/backlog"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")
	api := backlog.New(apiKey, baseURL, backlog.OptionDebug(true))

	fpath := filepath.Clean(filepath.Join("testdata", "test.jpg"))
	fileUploadResponse, err := api.UploadFile(fpath)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("id: %d, name: %s, size: %d\n", *fileUploadResponse.ID, *fileUploadResponse.Name, fileUploadResponse.Size)

	i := &backlog.AddAttachmentToWikiInput{
		WikiID:        backlog.Int(451845),
		AttachmentIDs: []int{*fileUploadResponse.ID},
	}
	attachments, err := api.AddAttachmentToWiki(i)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	log.Println("attachments", attachments)
}
