package backlog

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestErrorResponse(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusNotFound)
		if _, err := fmt.Fprint(w, `{"errors":[{"message": "No space.", "code": 6, "moreInfo": ""}]}`); err != nil {
			t.Fatal(err)
		}
	})

	client.debug = true
	client.httpclient = &http.Client{}
	client.log = log.New(os.Stdout, "backlog: ", log.Lshortfile|log.LstdFlags)

	client.Debugf("%s", "test")
	client.Debugln("test")

	if _, err := client.GetSpace(); err == nil {
		t.Fatal("expected an error but got none", err)
	}
}

func TestErrorResponseWithoutErrors(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusNotFound)
		if _, err := fmt.Fprint(w, `{"errors":[]}`); err != nil {
			t.Fatal(err)
		}
	})

	if _, err := client.GetSpace(); err == nil {
		t.Fatal("expected an error but got none", err)
	}
}

func TestStatusUnAuthorized(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
	})

	_, err := client.GetSpace()
	if err == nil {
		t.Fatal("expected an error but got none", err)
	}
}
