package backlog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

const testJSONWiki string = `{
	"id": 1,
	"projectId": 1,
	"name": "Home",
	"content": "test",
	"tags": [
		{
			"id": 12,
			"name": "議事録"
		}
	],
	"attachments": [
		{
			"id": 1,
			"name": "test.json",
			"size": 8857,
			"createdUser": {
				"id": 1,
				"userId": "admin",
				"name": null,
				"roleType": 1,
				"lang": "ja",
				"mailAddress": "eguchi@nulab.example"
			},
			"created": "2006-01-02T15:04:05Z"
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

func getTestWikiCount() Page {
	return Page{
		Count: Int(1),
	}
}

func getTestWiki() *Wiki {
	return &Wiki{
		ID:        Int(1),
		ProjectID: Int(1),
		Name:      String("Home"),
		Content:   String("test"),
		Tags: []*Tag{
			{
				ID:   Int(12),
				Name: String("議事録"),
			},
		},
		Attachments: []*Attachment{
			{
				ID:   Int(1),
				Name: String("test.json"),
				Size: Int(8857),
				CreatedUser: &User{
					ID:          Int(1),
					UserID:      String("admin"),
					RoleType:    RoleType(1),
					Lang:        String("ja"),
					MailAddress: String("eguchi@nulab.example"),
				},
				Created: &Timestamp{referenceTime},
			},
		},
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

func getTestWikiTag() *Tag {
	return &Tag{
		ID:   Int(1),
		Name: String("test"),
	}
}

func getTestGetWikiAttachment() *Attachment {
	return &Attachment{
		ID:   Int(1),
		Name: String("Duke.png"),
		Size: Int(196186),
	}
}

func getTestAddAttachmentToWiki() *Attachment {
	return &Attachment{
		ID:   Int(1),
		Name: String("Duke.png"),
		Size: Int(196186),
		CreatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    1,
			Lang:        nil,
			MailAddress: String("eguchi@nulab.example"),
		},
		Created: &Timestamp{referenceTime},
	}
}

func TestGetWikis(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONWiki)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	wikis, err := client.GetWikis(1, "test")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Wiki{getTestWiki()}
	if !reflect.DeepEqual(want, wikis) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, wikis)))
	}
}

func TestGetWikisWithInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetWikis(true, "test")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikisFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetWikis(1, "test"); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiCount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/count", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"count": 1}`); err != nil {
			t.Fatal(err)
		}
	})

	count, err := client.GetWikiCount(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWikiCount()
	if !reflect.DeepEqual(*want.Count, count) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetWikiCountFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/count", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetWikiCount(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiCountWithInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetWikiCount(true)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiTags(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/tags", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `
			[
				{
					"id": 1,
					"name": "test"
				}
			]
		`); err != nil {
			t.Fatal(err)
		}
	})

	tags, err := client.GetWikiTags(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Tag{getTestWikiTag()}
	if !reflect.DeepEqual(want, tags) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetWikiTagsWithInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetWikiTags(true)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiTagsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/tags", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetWikiTags(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONWiki); err != nil {
			t.Fatal(err)
		}
	})

	wiki, err := client.GetWiki(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWiki()
	if !reflect.DeepEqual(want, wiki) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetWikiByIDFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetWiki(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateWiki(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONWiki); err != nil {
			t.Fatal(err)
		}
	})

	input := &CreateWikiInput{
		ProjectID:  Int(1),
		Name:       String("Home"),
		Content:    String("test"),
		MailNotify: Bool(false),
	}
	wiki, err := client.CreateWiki(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWiki()
	if !reflect.DeepEqual(want, wiki) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateWikiFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &CreateWikiInput{
		ProjectID: Int(1),
		Name:      String("Home"),
		Content:   String("test"),
	}

	_, err := client.CreateWiki(input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateWiki(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONWiki); err != nil {
			t.Fatal(err)
		}
	})

	input := &UpdateWikiInput{
		WikiID:     Int(1),
		Name:       String("Home"),
		Content:    String("test"),
		MailNotify: Bool(false),
	}
	wiki, err := client.UpdateWiki(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWiki()
	if !reflect.DeepEqual(want, wiki) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateWikiFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &UpdateWikiInput{
		WikiID:  Int(1),
		Name:    String("Home"),
		Content: String("test"),
	}
	if _, err := client.UpdateWiki(input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteWiki(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONWiki); err != nil {
			t.Fatal(err)
		}
	})

	wiki, err := client.DeleteWiki(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWiki()
	if !reflect.DeepEqual(want, wiki) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteWikiFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.DeleteWiki(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetAttachmentToWiki(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1/attachments", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `
			[
				{
					"id": 1,
					"name": "Duke.png",
					"size": 196186
				}
			]
		`); err != nil {
			t.Fatal(err)
		}
	})

	attachments, err := client.GetWikiAttachments(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Attachment{getTestGetWikiAttachment()}
	if !reflect.DeepEqual(want, attachments) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetWikiAttachmentsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1/attachments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetWikiAttachments(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddAttachmentToWiki(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1/attachments", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `
		[
			{
				"id": 1,
				"name": "Duke.png",
				"size": 196186,
				"createdUser": {
					"id": 1,
					"userId": "admin",
					"name": "admin",
					"roleType": 1,
					"lang": null,
					"mailAddress": "eguchi@nulab.example"
				},
				"created": "2006-01-02T15:04:05Z"
			}
		]
		`); err != nil {
			t.Fatal(err)
		}
	})

	input := &AddAttachmentToWikiInput{
		WikiID:        Int(1),
		AttachmentIDs: []int{1},
	}
	attachment, err := client.AddAttachmentToWiki(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Attachment{getTestAddAttachmentToWiki()}
	if !reflect.DeepEqual(want, attachment) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestAddAttachmentToWikiFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1/attachments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &AddAttachmentToWikiInput{
		WikiID:        Int(1),
		AttachmentIDs: []int{1},
	}
	if _, err := client.AddAttachmentToWiki(input); err == nil {
		t.Fatal("expected an error but got none")
	}
}
