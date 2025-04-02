package backlog

import (
	"bytes"
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

var testJSONRecentlyViewedWiki = fmt.Sprintf(`{
    "page": %s,
    "updated": "2006-01-02T15:04:05Z"
}`, testJSONWiki)

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
		ID:          Int(1),
		Name:        String("Duke.png"),
		Size:        Int(196186),
		CreatedUser: getTestUser(),
		Created:     &Timestamp{referenceTime},
	}
}

func TestGetMyRecentlyViewedWikis(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/myself/recentlyViewedWikis", func(w http.ResponseWriter, _ *http.Request) {
		j := fmt.Sprintf(`[%s]`, testJSONRecentlyViewedWiki)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetMyRecentlyViewedWikis(&GetMyRecentlyViewedWikisOptions{
		Order:  Order(OrderAsc.String()),
		Offset: Int(10),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*RecentlyViewedWiki{
		{
			Page:    getTestWiki(),
			Updated: &Timestamp{referenceTime},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetMyRecentlyViewedWikisFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/myself/recentlyViewedWikis", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetMyRecentlyViewedWikis(&GetMyRecentlyViewedWikisOptions{
		Order:  Order(OrderAsc.String()),
		Offset: Int(10),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikis(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis", func(w http.ResponseWriter, _ *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONWiki)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	wikis, err := client.GetWikis(&GetWikisOptions{
		ProjectIDOrKey: 1,
		Keyword:        String("test"),
	})
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

	_, err := client.GetWikis(&GetWikisOptions{
		ProjectIDOrKey: true,
		Keyword:        String("test"),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikisFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetWikis(&GetWikisOptions{
		ProjectIDOrKey: 1,
		Keyword:        String("test"),
	}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiCount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/count", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, `{"count": 1}`); err != nil {
			t.Fatal(err)
		}
	})

	count, err := client.GetWikiCount(&GetWikiCountOptions{})
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

	mux.HandleFunc("/wikis/count", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetWikiCount(&GetWikiCountOptions{
		ProjectIDOrKey: 1,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiCountWithInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetWikiCount(&GetWikiCountOptions{
		ProjectIDOrKey: true,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiTags(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/tags", func(w http.ResponseWriter, _ *http.Request) {
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

	tags, err := client.GetWikiTags(&GetWikiTagsOptions{})
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

	_, err := client.GetWikiTags(&GetWikiTagsOptions{
		ProjectIDOrKey: true,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiTagsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/tags", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetWikiTags(&GetWikiTagsOptions{
		ProjectIDOrKey: 1,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, _ *http.Request) {
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

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, _ *http.Request) {
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

	mux.HandleFunc("/wikis", func(w http.ResponseWriter, _ *http.Request) {
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

	mux.HandleFunc("/wikis", func(w http.ResponseWriter, _ *http.Request) {
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

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONWiki); err != nil {
			t.Fatal(err)
		}
	})

	input := &UpdateWikiInput{
		Name:       String("Home"),
		Content:    String("test"),
		MailNotify: Bool(false),
	}
	wiki, err := client.UpdateWiki(1, input)
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

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &UpdateWikiInput{
		Name:    String("Home"),
		Content: String("test"),
	}
	if _, err := client.UpdateWiki(1, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteWiki(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, _ *http.Request) {
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

	mux.HandleFunc("/wikis/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.DeleteWiki(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetAttachmentToWiki(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1/attachments", func(w http.ResponseWriter, _ *http.Request) {
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

	mux.HandleFunc("/wikis/1/attachments", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetWikiAttachments(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiAttachmentContent(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	client.httpclient = &mockHTTPClient{}

	err := client.GetWikiAttachmentContent(1, 1, &bytes.Buffer{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestGetWikiAttachmentContentFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1/attachments/2", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.GetWikiAttachmentContent(1, 2, &bytes.Buffer{}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddAttachmentToWiki(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1/attachments", func(w http.ResponseWriter, _ *http.Request) {
		j := fmt.Sprintf(`[%s]`, testJSONAttachment)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	input := &AddAttachmentToWikiInput{
		AttachmentIDs: []int{1},
	}
	attachment, err := client.AddAttachmentToWiki(1, input)
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

	mux.HandleFunc("/wikis/1/attachments", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &AddAttachmentToWikiInput{
		AttachmentIDs: []int{1},
	}
	if _, err := client.AddAttachmentToWiki(1, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteWikiAttachment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1/attachments/8", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONAttachment); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteAttachmentInWiki(1, 8)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestAddAttachmentToWiki()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteAttachmentInWikiFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wikis/1/attachments/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.DeleteAttachmentInWiki(1, 1); err == nil {
		t.Fatal("expected an error but got none")
	}
}
