package backlog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

const testJSONCustomField string = `{
	"id": 1,
	"typeId": 6,
	"name": "custom",
	"description": "",
	"required": false,
	"applicableIssueTypes": [],
	"allowAddItem": false,
	"items": [
		{
			"id": 1,
			"name": "Windows 8",
			"displayOrder": 0
		}
	]
}`

func getTestCustomField() *CustomField {
	return &CustomField{
		ID:                   Int(1),
		TypeID:               Int(6),
		Name:                 String("custom"),
		Description:          String(""),
		Required:             Bool(false),
		ApplicableIssueTypes: []int{},
		AllowAddItem:         Bool(false),
		Items: []*Item{
			{
				ID:           Int(1),
				Name:         String("Windows 8"),
				DisplayOrder: Int(0),
			},
		},
	}
}

func TestGetCustomFields(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/customFields", func(w http.ResponseWriter, _ *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONCustomField)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetCustomFields("SRE")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*CustomField{getTestCustomField()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetCustomFieldsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/customFields", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetCustomFields("SRE")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetCustomFieldsInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetCustomFields("%%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
