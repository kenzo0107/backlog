package backlog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
)

func getTestUploadFile() FileUploadResponse {
	return getTestUploadFileWithID(1)
}

func getTestUploadFileWithID(id int) FileUploadResponse {
	return FileUploadResponse{
		ID:   id,
		Name: "test.txt",
		Size: 8857,
	}
}

func getUploadFile(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestUploadFile())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func TestUploadFile(t *testing.T) {
	http.HandleFunc("/api/v2/space/attachment", getUploadFile)
	expected := getTestUploadFile()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	fpath := filepath.Clean(filepath.Join("testdata", "test.jpg"))
	fileUploadResponse, err := api.UploadFile(fpath)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, *fileUploadResponse) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUploadFileFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/wikis", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	fpath := filepath.Clean(filepath.Join("testdata", "test.jpg"))

	if _, err := api.UploadFile(fpath); err == nil {
		t.Fatal("expected an error but got none")
	}
}
