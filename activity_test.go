package backlog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

const testJSONActivity string = `{
	"id": 1,
	"project": {
		"id": 92,
		"projectKey": "SUB",
		"name": "サブタスク",
		"chartEnabled": true,
		"subtaskingEnabled": true,
		"projectLeaderCanEditProjectLeader": false,
		"textFormattingRule": "",
		"archived": false,
		"displayOrder": 0
	},
	"type": 2,
	"content": {
		"id": 4809,
		"key_id": 121,
		"summary": "コメント",
		"description": "",
		"comment": {
			"id": 7237,
			"content": ""
		},
		"changes": [
			{
				"field": "milestone",
				"new_value": " R2014-07-23",
				"old_value": "",
				"type": "standard"
			},
			{
				"field": "status",
				"new_value": "4",
				"old_value": "1",
				"type": "standard"
			}
		]
	},
	"notifications": [
		{
			"id": 25,
			"alreadyRead": false,
			"reason": 2,
			"user": {
				"id": 5686,
				"userId": "takada",
				"name": "takada",
				"roleType": 2,
				"lang": "ja",
				"mailAddress": "takada@nulab.example"
			},
			"resourceAlreadyRead": false
		}
	],
	"createdUser": {
		"id": 1,
		"userId": "admin",
		"name": "admin",
		"roleType": 1,
		"lang": "ja",
		"mailAddress": "eguchi@nulab.example"
	},
	"created": "2006-01-02T15:04:05Z"
}`

func TestGetUserActivities(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1/activities", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONActivity)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	input := &GetUserActivitiesOptions{
		ActivityTypeIDs: []int{1, 2, 3},
		MinID:           Int(1),
		MaxID:           Int(10),
		Count:           Int(20),
		Order:           OrderAsc,
	}
	activities, err := client.GetUserActivities(1, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Activity{getTestActivity()}

	if !reflect.DeepEqual(want, activities) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserActivitiesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1/activities", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &GetUserActivitiesOptions{}
	if _, err := client.GetUserActivities(1, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectActivities(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/activities", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONActivity)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetProjectActivities("SRE", &GetProjectActivitiesOptions{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Activity{getTestActivity()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetProjectActivitiesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/activities", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetProjectActivities("SRE", &GetProjectActivitiesOptions{}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectActivitiesInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetProjectActivities("%", &GetProjectActivitiesOptions{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
