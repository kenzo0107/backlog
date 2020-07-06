package backlog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const testJSONWebhook string = `{
	"id": 1,
	"name": "webhook",
	"description": "",
	"hookUrl": "http://nulab.test/",
	"allEvent": false,
	"activityTypeIds": [1, 2, 3, 4, 5],
	"createdUser": {
		"id": 1,
		"userId": "admin",
		"name": "admin",
		"roleType": 1,
		"lang": "ja",
		"mailAddress": "eguchi@nulab.example"
	},
	"created": "2006-01-02T15:04:05Z",
	"updatedUser": {
		"id": 1,
		"userId": "admin",
		"name": "admin",
		"roleType": 1,
		"lang": "ja",
		"mailAddress": "eguchi@nulab.example"
	},
	"updated": "2006-01-02T15:04:05Z"
}`

func getTestWebhook() *Webhook {
	return &Webhook{
		ID:              Int(1),
		Name:            String("webhook"),
		Description:     String(""),
		HookURL:         String("http://nulab.test/"),
		AllEvent:        Bool(false),
		ActivityTypeIds: []int{1, 2, 3, 4, 5},
		CreatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    RoleType(1),
			Lang:        String("ja"),
			MailAddress: String("eguchi@nulab.example"),
		},
		Created: &Timestamp{referenceTime},
		UpdatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    RoleType(1),
			Lang:        String("ja"),
			MailAddress: String("eguchi@nulab.example"),
		},
		Updated: &Timestamp{referenceTime},
	}
}

func TestGetWebhooks(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1/webhooks", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONWebhook)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	webhooks, err := client.GetWebhooks(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Webhook{getTestWebhook()}
	if !reflect.DeepEqual(want, webhooks) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetWebhooksWithInvalidProjectID(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if _, err := client.GetWebhooks(false); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWebhooksFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1/webhooks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetWebhooks(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/webhooks", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONWebhook); err != nil {
			t.Fatal(err)
		}
	})

	input := &CreateWebhookInput{
		Name:            String("webhook"),
		Description:     String(""),
		HookURL:         String("https://webhook.example.com"),
		AllEvent:        Bool(false),
		ActivityTypeIDs: []int{1, 2, 3, 4, 5},
	}
	webhook, err := client.CreateWebhook("SRE", input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWebhook()
	if !reflect.DeepEqual(want, webhook) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestAddWebhookWithInvalidProjectID(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	input := &CreateWebhookInput{
		Name:    String("webhook"),
		HookURL: String("https://webhook.example.com"),
	}
	if _, err := client.CreateWebhook(true, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddWebhookFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/SRE/webhooks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &CreateWebhookInput{
		Name:    String("webhook"),
		HookURL: String("https://webhook.example.com"),
	}
	if _, err := client.CreateWebhook("SRE", input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/webhooks/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONWebhook); err != nil {
			t.Fatal(err)
		}
	})

	webhook, err := client.GetWebhook("SRE", 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWebhook()
	if !reflect.DeepEqual(want, webhook) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetWebhookWithInvalidProjectID(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if _, err := client.GetWebhook(true, 1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWebhookFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/SRE/webhooks/10", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetWebhook("SRE", 10); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/webhooks/10", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONWebhook); err != nil {
			t.Fatal(err)
		}
	})

	input := &UpdateWebhookInput{
		Name:            String("webhook"),
		Description:     String(""),
		HookURL:         String("https://webhook.example.com"),
		AllEvent:        Bool(false),
		ActivityTypeIDs: []int{1, 2, 3, 4, 5},
	}
	webhook, err := client.UpdateWebhook("SRE", 10, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWebhook()
	if !reflect.DeepEqual(want, webhook) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateWebhookWithInvalidProjectID(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	input := &UpdateWebhookInput{
		Name:    String("webhook"),
		HookURL: String("https://webhook.example.com"),
	}
	if _, err := client.UpdateWebhook("SRE", 10, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateWebhookFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/SRE/webhooks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &UpdateWebhookInput{
		Name:    String("webhook"),
		HookURL: String("https://webhook.example.com"),
	}
	if _, err := client.UpdateWebhook("SRE", 1, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/webhooks/10", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONWebhook); err != nil {
			t.Fatal(err)
		}
	})

	webhook, err := client.DeleteWebhook("SRE", 10)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWebhook()
	if !reflect.DeepEqual(want, webhook) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteWebhookWithInvalidProjectID(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if _, err := client.DeleteWebhook(true, 10); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteWebhookFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/SRE/webhooks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.DeleteWebhook("SRE", 10); err == nil {
		t.Fatal("expected an error but got none")
	}
}
