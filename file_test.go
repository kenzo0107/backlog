package backlog

import (
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
)

func getTestUploadFile() *FileUploadResponse {
	return &FileUploadResponse{
		ID:   Int(1),
		Name: String("test.txt"),
		Size: Int(8857),
	}
}

func TestUploadFile(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/attachment", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, `{"id": 1, "name": "test.txt", "size": 8857}`); err != nil {
			t.Fatal(err)
		}
	})

	fpath := filepath.Clean(filepath.Join("testdata", "test.jpg"))
	fileUploadResponse, err := client.UploadFile(fpath)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUploadFile()
	if !reflect.DeepEqual(want, fileUploadResponse) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUploadFileFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/attachment", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	fpath := filepath.Clean(filepath.Join("testdata", "test.jpg"))
	if _, err := client.UploadFile(fpath); err == nil {
		t.Fatal("expected an error but got none")
	}
}
