package backlog

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

const testJSONSpace string = `{
	"spaceKey": "nulab",
	"name": "Nulab Inc.",
	"ownerId": 1,
	"lang": "ja",
	"timezone": "Asia/Tokyo",
	"reportSendTime": "08:00:00",
	"textFormattingRule": "markdown",
	"created": "2006-01-02T15:04:05Z",
	"updated": "2006-01-02T15:04:05Z"
}`

func getTestSpace() *Space {
	return &Space{
		SpaceKey:           String("nulab"),
		Name:               String("Nulab Inc."),
		OwnerID:            Int(1),
		Lang:               String("ja"),
		Timezone:           String("Asia/Tokyo"),
		ReportSendTime:     String("08:00:00"),
		TextFormattingRule: String("markdown"),
		Created:            &Timestamp{referenceTime},
		Updated:            &Timestamp{referenceTime},
	}
}

func getTestSpaceNotification() *SpaceNotification {
	return &SpaceNotification{
		Content: String("Notification"),
		Updated: &Timestamp{referenceTime},
	}
}

func getTestSpaceDiskUsage() *SpaceDiskUsage {
	return &SpaceDiskUsage{
		Capacity:   Int(1073741824),
		Issue:      Int(119511),
		Wiki:       Int(48575),
		File:       Int(0),
		Subversion: Int(0),
		Git:        Int(0),
		GitLFS:     Int(0),
		Details: []*SpaceDiskUsageDetail{
			{
				ProjectID:  Int(1),
				Issue:      Int(11931),
				Wiki:       Int(0),
				File:       Int(0),
				Subversion: Int(0),
				Git:        Int(0),
				GitLFS:     Int(0),
			},
		},
	}
}

func TestGetSpace(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, testJSONSpace); err != nil {
			t.Fatal(err)
		}
	})

	space, err := client.GetSpace()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestSpace()
	if !reflect.DeepEqual(want, space) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSpaceFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetSpace(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

type mockHTTPClient struct{}

func (m *mockHTTPClient) Do(*http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(`OK`))}, nil
}

func TestGetSpaceIcon(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	client.httpclient = &mockHTTPClient{}

	err := client.GetSpaceIcon(&bytes.Buffer{})
	if err != nil {
		log.Fatalf("Unexpected error: %s in test", err)
	}
}

func TestGetSpaceIconFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/image", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.GetSpaceIcon(&bytes.Buffer{}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSpaceNotification(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/notification", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"content": "Notification", "updated": "2006-01-02T15:04:05Z"}`); err != nil {
			t.Fatal(err)
		}
	})

	spaceNotification, err := client.GetSpaceNotification()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestSpaceNotification()
	if !reflect.DeepEqual(want, spaceNotification) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSpaceNotificationFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/notification", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetSpaceNotification(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateSpaceNotification(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/notification", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"content": "Notification", "updated": "2006-01-02T15:04:05Z"}`); err != nil {
			t.Fatal(err)
		}
	})

	spaceNotification, err := client.UpdateSpaceNotification(&UpdateSpaceNotificationInput{
		Content: String("test"),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestSpaceNotification()
	if !reflect.DeepEqual(want, spaceNotification) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateSpaceNotificationFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/notification", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.UpdateSpaceNotification(&UpdateSpaceNotificationInput{
		Content: String("test"),
	}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSpaceDiskUsage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/diskUsage", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `
			{
				"capacity": 1073741824,
				"issue": 119511,
				"wiki": 48575,
				"file": 0,
				"subversion": 0,
				"git": 0,
				"gitLFS": 0,
				"details":[
					{
						"projectId": 1,
						"issue": 11931,
						"wiki": 0,
						"file": 0,
						"subversion": 0,
						"git": 0,
						"gitLFS": 0
					}
				]
			}
		`); err != nil {
			t.Fatal(err)
		}
	})

	spaceNotification, err := client.GetSpaceDiskUsage()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestSpaceDiskUsage()
	if !reflect.DeepEqual(want, spaceNotification) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSpaceDiskUsageFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/diskUsage", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetSpaceDiskUsage(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLicence(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/licence", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"active": true,
			"attachmentLimit": 0,
			"attachmentLimitPerFile": 10485760,
			"attachmentNumLimit": 50,
			"attribute": true,
			"attributeLimit": 100,
			"burndown": true,
			"commentLimit": 0,
			"componentLimit": 0,
			"fileSharing": true,
			"gantt": true,
			"git": true,
			"issueLimit": 0,
			"licenceTypeId": 51,
			"limitDate": "2006-01-02T15:04:05Z",
			"nulabAccount": true,
			"parentChildIssue": true,
			"postIssueByMail": true,
			"projectGroup": true,
			"projectLimit": 0,
			"pullRequestAttachmentLimitPerFile": 10485760,
			"pullRequestAttachmentNumLimit": 50,
			"remoteAddress": true,
			"remoteAddressLimit": 100,
			"startedOn": "2006-01-02T15:04:05Z",
			"storageLimit": 1073741824000,
			"subversion": true,
			"subversionExternal": true,
			"userLimit": 0,
			"versionLimit": 0,
			"wikiAttachment": true,
			"wikiAttachmentLimitPerFile": 10485760,
			"wikiAttachmentNumLimit": 50
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLicence()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &License{
		Active:                            Bool(true),
		AttachmentLimit:                   Int(0),
		AttachmentLimitPerFile:            Int(10485760),
		AttachmentNumLimit:                Int(50),
		Attribute:                         Bool(true),
		AttributeLimit:                    Int(100),
		Burndown:                          Bool(true),
		CommentLimit:                      Int(0),
		ComponentLimit:                    Int(0),
		FileSharing:                       Bool(true),
		Gantt:                             Bool(true),
		Git:                               Bool(true),
		IssueLimit:                        Int(0),
		LicenceTypeID:                     Int(51),
		LimitDate:                         &Timestamp{referenceTime},
		NulabAccount:                      Bool(true),
		ParentChildIssue:                  Bool(true),
		PostIssueByMail:                   Bool(true),
		ProjectGroup:                      Bool(true),
		ProjectLimit:                      Int(0),
		PullRequestAttachmentLimitPerFile: Int(10485760),
		PullRequestAttachmentNumLimit:     Int(50),
		RemoteAddress:                     Bool(true),
		RemoteAddressLimit:                Int(100),
		StartedOn:                         &Timestamp{referenceTime},
		StorageLimit:                      Int64(1073741824000),
		Subversion:                        Bool(true),
		SubversionExternal:                Bool(true),
		UserLimit:                         Int(0),
		VersionLimit:                      Int(0),
		WikiAttachment:                    Bool(true),
		WikiAttachmentLimitPerFile:        Int(10485760),
		WikiAttachmentNumLimit:            Int(50),
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSpaceLicenseFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/licence", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetLicence(); err == nil {
		t.Fatal("expected an error but got none")
	}
}
