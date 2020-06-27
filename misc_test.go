package backlog

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestProjectIDOrKey(t *testing.T) {
	type args struct {
		projectIDOrKey interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "set projectIDOrKey int",
			args: args{
				projectIDOrKey: 12345,
			},
			want:    "12345",
			wantErr: false,
		},
		{
			name: "set projectIDOrKey string",
			args: args{
				projectIDOrKey: "AIUEO",
			},
			want:    "AIUEO",
			wantErr: false,
		},
		{
			name: "set projectIDOrKey bool",
			args: args{
				projectIDOrKey: true,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "set projectIDOrKey []string",
			args: args{
				projectIDOrKey: []string{"a"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := projIDOrKey(tt.args.projectIDOrKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectIDOrKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("projectIDOrKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

	err.Error()
	// err.HTTPStatusCode()
}
