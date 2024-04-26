package backlog

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/pkg/errors"

	"github.com/kylelemons/godebug/pretty"
)

var testJSONSharedFile = fmt.Sprintf(`{
	"id": 454403,
	"type": "file",
	"dir": "/ユーザアイコン/",
	"name": "01_サラリーマン.png",
	"size": 2735,
	"createdUser": %s,
	"created": "2006-01-02T15:04:05Z",
	"updatedUser": %s,
	"updated": "2006-01-02T15:04:05Z"
}`, testJSONUser, testJSONUser)

var testJSONIssue = fmt.Sprintf(`{
	"id": 1,
	"projectId": 1,
	"issueKey": "BLG-1",
	"keyId": 1,
	"issueType": {
		"id": 2,
		"projectId": 1,
		"name": "タスク",
		"color": "#7ea800",
		"displayOrder": 0
	},
	"summary": "first issue",
	"description": "",
	"resolution": {
		"id": 0,
		"name": "対応済み"
	},
	"priority": {
		"id": 3,
		"name": "中"
	},
	"status": {
		"id": 1,
		"projectId": 1,
		"name": "未対応",
		"color": "#ed8077",
		"displayOrder": 1000
	},
	"assignee": %s,
	"category": [],
	"versions": [],
	"milestone": [
		{
			"id": 30,
			"projectId": 1,
			"name": "wait for release",
			"description": "",
			"startDate": null,
			"releaseDueDate": null,
			"archived": false
		}
	],
	"startDate": null,
	"dueDate": null,
	"estimatedHours": null,
	"actualHours": null,
	"parentIssueId": null,
	"createdUser": %s,
	"created": "2006-01-02T15:04:05Z",
	"updatedUser": %s,
	"updated": "2006-01-02T15:04:05Z",
	"customFields": [
		{
			"id": 111,
			"fieldTypeId": 1,
			"name": "custom string",
			"value": "hoge"
		}
	],
	"attachments": [
		{
			"id": 1,
			"name": "IMGP0088.JPG",
			"size": 85079
		}
	],
	"sharedFiles": [%s],
	"stars": [
		{
			"id": 10,
			"comment": null,
			"url": "https://xx.backlog.jp/view/BLG-1",
			"title": "[BLG-1] first issue | 課題の表示 - Backlog",
			"presenter": %s,
			"created": "2006-01-02T15:04:05Z"
		}
	]
}`, testJSONUser, testJSONUser, testJSONUser, testJSONSharedFile, testJSONUser)

var testJSONNotification = fmt.Sprintf(`{
	"id":22,
	"alreadyRead":false,
	"reason":2,
	"user": %s,
	"resourceAlreadyRead":false
}`, testJSONUser)

var testJSONIssueComment = fmt.Sprintf(`{
    "id": 6586,
    "content": "テスト",
    "changeLog": null,
    "createdUser": %s,
    "created": "2006-01-02T15:04:05Z",
    "updated": "2006-01-02T15:04:05Z",
    "stars": [],
    "notifications": [%s]
}`, testJSONUser, testJSONNotification)

var testJSONAttachment = fmt.Sprintf(`{
	"id": 1,
	"name": "Duke.png",
	"size": 196186,
	"createdUser": %v,
	"created": "2006-01-02T15:04:05Z"
}`, testJSONUser)

func getTestIssue() *Issue {
	return &Issue{
		ID:        Int(1),
		ProjectID: Int(1),
		IssueKey:  String("BLG-1"),
		KeyID:     Int(1),
		IssueType: &IssueType{
			ID:           Int(2),
			ProjectID:    Int(1),
			Name:         String("タスク"),
			Color:        String("#7ea800"),
			DisplayOrder: Int(0),
		},
		Summary:     String("first issue"),
		Description: String(""),
		Resolution: &Resolution{
			ID:   Int(0),
			Name: String("対応済み"),
		},
		Priority: &Priority{
			ID:   Int(3),
			Name: String("中"),
		},
		Status: &Status{
			ID:           Int(1),
			ProjectID:    Int(1),
			Name:         String("未対応"),
			Color:        String("#ed8077"),
			DisplayOrder: Int(1000),
		},
		Assignee: getTestUser(),
		Category: []*Category{},
		Versions: []*Version{},
		Milestone: []*Milestone{
			{
				ID:             Int(30),
				ProjectID:      Int(1),
				Name:           String("wait for release"),
				Description:    String(""),
				StartDate:      nil,
				ReleaseDueDate: nil,
				Archived:       Bool(false),
			},
		},
		StartDate:      nil,
		DueDate:        nil,
		EstimatedHours: nil,
		ActualHours:    nil,
		ParentIssueID:  nil,
		CreatedUser:    getTestUser(),
		Created:        &Timestamp{referenceTime},
		UpdatedUser:    getTestUser(),
		Updated:        &Timestamp{referenceTime},
		CustomFields: []*IssueCustomField{
			{
				ID:          Int(111),
				FieldTypeID: Int(1),
				Name:        String("custom string"),
				Value:       "hoge",
			},
		},
		Attachments: []*Attachment{
			{
				ID:   Int(1),
				Name: String("IMGP0088.JPG"),
				Size: Int(85079),
			},
		},
		SharedFiles: []*SharedFile{getTestSharedFile()},
		Stars: []*Star{
			{
				ID:        Int(10),
				Comment:   nil,
				URL:       String("https://xx.backlog.jp/view/BLG-1"),
				Title:     String("[BLG-1] first issue | 課題の表示 - Backlog"),
				Presenter: getTestUser(),
				Created:   &Timestamp{referenceTime},
			},
		},
	}
}

func getTestIssueComment() *IssueComment {
	return &IssueComment{
		ID:          Int(6586),
		Content:     String("テスト"),
		ChangeLog:   nil,
		CreatedUser: getTestUser(),
		Created:     &Timestamp{referenceTime},
		Updated:     &Timestamp{referenceTime},
		Stars:       []*Star{},
		Notifications: []*Notification{
			getTestNotification(),
		},
	}
}

func getTestNotification() *Notification {
	return &Notification{
		ID:                  Int(22),
		AlreadyRead:         Bool(false),
		Reason:              Int(2),
		User:                getTestUser(),
		ResourceAlreadyRead: Bool(false),
	}
}

func getTestAttachment() *Attachment {
	return &Attachment{
		ID:          Int(1),
		Name:        String("Duke.png"),
		Size:        Int(196186),
		CreatedUser: getTestUser(),
		Created:     &Timestamp{referenceTime},
	}
}

func getTestSharedFile() *SharedFile {
	return &SharedFile{
		ID:          Int(454403),
		Type:        String("file"),
		Dir:         String("/ユーザアイコン/"),
		Name:        String("01_サラリーマン.png"),
		Size:        Int(2735),
		CreatedUser: getTestUser(),
		Created:     &Timestamp{referenceTime},
		UpdatedUser: getTestUser(),
		Updated:     &Timestamp{referenceTime},
	}
}

func TestGetIssues(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONIssue)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	opts := &GetIssuesOptions{
		ProjectIDs:     []int{1},
		IssueTypeIDs:   []int{2},
		CategoryIDs:    []int{3},
		VersionIDs:     []int{4},
		MilestoneIDs:   []int{5},
		StatusIDs:      []int{6},
		PriorityIDs:    []int{7},
		CreatedUserIDs: []int{8},
		ResolutionIDs:  []int{9},
		AssigneeIDs:    []int{10},
		ParentChild:    Int(11),
		Attachment:     Bool(false),
		SharedFile:     Bool(false),
		Sort:           SortIssueType,
		Offset:         Int(10),
		CreatedSince:   String("2019-01-07"),
		CreatedUntil:   String("2020-01-07"),
		UpdatedSince:   String("2019-01-07"),
		UpdatedUntil:   String("2020-01-07"),
		StartDateSince: String("2019-01-07"),
		StartDateUntil: String("2020-01-07"),
		DueDateSince:   String("2019-01-07"),
		DueDateUntil:   String("2020-01-07"),
		IDs:            []int{1},
		ParentIssueIDs: []int{11, 12, 13},
		Keyword:        String("test"),
	}
	issues, err := client.GetIssues(opts)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Issue{getTestIssue()}
	if !reflect.DeepEqual(want, issues) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetIssuesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	opts := &GetIssuesOptions{
		Sort: SortVersion,
	}
	if _, err := client.GetIssues(opts); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestSort_String(t *testing.T) {
	tests := []struct {
		name string
		k    Sort
		want string
	}{
		{
			name: "issueType",
			k:    SortIssueType,
			want: "issueType",
		},
		{
			name: "category",
			k:    SortCategory,
			want: "category",
		},
		{
			name: "version",
			k:    SortVersion,
			want: "version",
		},
		{
			name: "milestone",
			k:    SortMilestone,
			want: "milestone",
		},
		{
			name: "summary",
			k:    SortSummary,
			want: "summary",
		},
		{
			name: "status",
			k:    SortStatus,
			want: "status",
		},
		{
			name: "priority",
			k:    SortPriority,
			want: "priority",
		},
		{
			name: "attachment",
			k:    SortAttachment,
			want: "attachment",
		},
		{
			name: "sharedFile",
			k:    SortSharedFile,
			want: "sharedFile",
		},
		{
			name: "created",
			k:    SortCreated,
			want: "created",
		},
		{
			name: "createdUser",
			k:    SortCreatedUser,
			want: "createdUser",
		},
		{
			name: "updated",
			k:    SortUpdated,
			want: "updated",
		},
		{
			name: "updatedUser",
			k:    SortUpdatedUser,
			want: "updatedUser",
		},
		{
			name: "assignee",
			k:    SortAssignee,
			want: "assignee",
		},
		{
			name: "startDate",
			k:    SortStartDate,
			want: "startDate",
		},
		{
			name: "dueDate",
			k:    SortDueDate,
			want: "dueDate",
		},
		{
			name: "estimatedHours",
			k:    SortEstimatedHours,
			want: "estimatedHours",
		},
		{
			name: "actualHours",
			k:    SortActualHours,
			want: "actualHours",
		},
		{
			name: "childIssue",
			k:    SortChildIssue,
			want: "childIssue",
		},
		{
			name: "empty",
			k:    Sort(""),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("Sort.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserMySelfRecentrlyViewedIssues(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/myself/recentlyViewedIssues", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf(`[{"issue":%s, "updated": "2006-01-02T15:04:05Z"}]`, testJSONIssue)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	opts := &GetUserMySelfRecentrlyViewedIssuesOptions{
		Order:  OrderAsc,
		Offset: Int(1),
		Count:  Int(100),
	}
	issues, err := client.GetUserMySelfRecentrlyViewedIssues(opts)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := Issues{
		{getTestIssue()},
	}

	if !reflect.DeepEqual(want, issues) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, issues)))
	}
}

func TestGetUserMySelfRecentrlyViewedIssuesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/myself/recentlyViewedIssues", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	opts := &GetUserMySelfRecentrlyViewedIssuesOptions{}
	_, err := client.GetUserMySelfRecentrlyViewedIssues(opts)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserMySelfRecentrlyViewedIssues4xxErrorFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/myself/recentlyViewedIssues", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `
			{
				status:     "400 Bad Request",
				status_code: 400,
				proto:      "HTTP/2.0"
			}
		`); err != nil {
			t.Fatal(err)
		}
	})

	opts := &GetUserMySelfRecentrlyViewedIssuesOptions{}
	if _, err := client.GetUserMySelfRecentrlyViewedIssues(opts); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueCount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/count", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"count":1}`); err != nil {
			t.Fatal(err)
		}
	})

	opts := &GetIssuesCountOptions{
		Order:  OrderAsc,
		Offset: Int(1),
		Count:  Int(100),
	}
	expected, err := client.GetIssueCount(opts)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := 1
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetIssueCountFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/count", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetIssueCount(&GetIssuesCountOptions{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateIssue(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssue); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateIssue(&CreateIssueInput{
		ProjectID:      Int(1),
		Summary:        String("first issue"),
		ParentIssueID:  nil,
		Description:    String(""),
		StartDate:      nil,
		DueDate:        nil,
		EstimatedHours: nil,
		ActualHours:    nil,
		IssueTypeID:    Int(2),
		MilestoneIDs:   []int{30},
		PriorityID:     Int(3),
		AssigneeID:     Int(2),
		AttachmentIDs:  []int{1},
		CustomFields: []*IssueCustomField{
			{
				ID:          Int(111),
				FieldTypeID: Int(1),
				Name:        String("custom string"),
				Value:       "hoge",
			},
			{
				ID:          Int(222),
				FieldTypeID: Int(6),
				Name:        String("custom list"),
				Value: []*Item{
					{
						ID:           Int(1),
						Name:         String("item foo"),
						DisplayOrder: Int(0),
					},
					{
						ID:           Int(2),
						Name:         String("item bar"),
						DisplayOrder: Int(1),
					},
				},
			},
			{
				ID:          Int(333),
				FieldTypeID: Int(6),
				Name:        String("custom list 2"),
				Value:       []*Item{},
			},
		},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssue()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateIssueFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateIssue(&CreateIssueInput{
		ProjectID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssue(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssue); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetIssue("BLG-1")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssue()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetIssueFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetIssue("BLG-1")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetIssue("%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateIssue(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssue); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateIssue("BLG-1", &UpdateIssueInput{
		Summary:        String("first issue"),
		EstimatedHours: String(""), // set empty value
		CustomFields: []*IssueCustomField{
			{
				ID:          Int(111),
				FieldTypeID: Int(1),
				Name:        String("custom string"),
				Value:       "hoge",
			},
			{
				ID:          Int(222),
				FieldTypeID: Int(6),
				Name:        String("custom list"),
				Value: []*Item{
					{
						ID:           Int(1),
						Name:         String("item foo"),
						DisplayOrder: Int(0),
					},
					{
						ID:           Int(2),
						Name:         String("item bar"),
						DisplayOrder: Int(1),
					},
				},
			},
			{
				ID:          Int(333),
				FieldTypeID: Int(6),
				Name:        String("custom list 2"),
				Value:       []*Item{},
			},
		},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssue()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateIssueFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateIssue("BLG-1", &UpdateIssueInput{
		Summary: String("first issue"),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateIssueInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.UpdateIssue("%", &UpdateIssueInput{
		Summary: String("first issue"),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteIssue(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssue); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteIssue("BLG-1")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssue()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestDeleteIssueFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteIssue("BLG-1")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteIssueInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.DeleteIssue("%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueComments(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf(`[%s]`, testJSONIssueComment)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetIssueComments("BLG-1", &GetIssueCommentsOptions{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*IssueComment{getTestIssueComment()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetIssueCommentsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetIssueComments("BLG-1", &GetIssueCommentsOptions{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueCommentsInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetIssueComments("%", &GetIssueCommentsOptions{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateIssueComment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssueComment); err != nil {
			t.Fatal(err)
		}
	})

	input := &CreateIssueCommentInput{
		Content: String("テスト"),
	}
	expected, err := client.CreateIssueComment("BLG-1", input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssueComment()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateIssueCommentFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &CreateIssueCommentInput{
		Content: String("テスト"),
	}
	_, err := client.CreateIssueComment("BLG-1", input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateIssueCommentInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	input := &CreateIssueCommentInput{
		Content: String("テスト"),
	}
	_, err := client.CreateIssueComment("%", input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueCommentsCount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/count", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"count": 1}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetIssueCommentsCount("BLG-1")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(1, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetIssueCommentsCountFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/count", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetIssueCommentsCount("BLG-1")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueCommentsCountInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetIssueCommentsCount("%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueComment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssueComment); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetIssueComment("BLG-1", 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssueComment()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetIssueCommentFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetIssueComment("BLG-1", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueCommentInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetIssueComment("%", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteIssueComment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssueComment); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteIssueComment("BLG-1", 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssueComment()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteIssueCommentFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteIssueComment("BLG-1", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteIssueCommentInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.DeleteIssueComment("%", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateIssueComment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssueComment); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateIssueComment("BLG-1", 1, &UpdateIssueCommentInput{
		Content: String("テスト"),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssueComment()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateIssueCommentFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateIssueComment("BLG-1", 1, &UpdateIssueCommentInput{
		Content: String("テスト"),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateIssueCommentInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.UpdateIssueComment("%", 1, &UpdateIssueCommentInput{
		Content: String("テスト"),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueCommentsNotifications(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/1/notifications", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf(`[%s]`, testJSONNotification)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetIssueCommentsNotifications("BLG-1", 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Notification{getTestNotification()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetIssueCommentsNotificationsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/1/notifications", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetIssueCommentsNotifications("BLG-1", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueCommentsNotificationsInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetIssueCommentsNotifications("%", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateIssueCommentsNotification(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/1/notifications", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssueComment); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateIssueCommentsNotification("BLG-1", 1, &CreateIssueCommentsNotificationInput{
		NotifiedUserIDs: []int{1},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssueComment()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateIssueCommentsNotificationFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/comments/1/notifications", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateIssueCommentsNotification("BLG-1", 1, &CreateIssueCommentsNotificationInput{
		NotifiedUserIDs: []int{1},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateIssueCommentsNotificationInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.CreateIssueCommentsNotification("%", 1, &CreateIssueCommentsNotificationInput{
		NotifiedUserIDs: []int{1},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueAttachments(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/attachments", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf(`[%s]`, testJSONAttachment)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetIssueAttachments("BLG-1")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Attachment{getTestAttachment()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetIssueAttachmentsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/attachments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetIssueAttachments("BLG-1")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueAttachmentsInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetIssueAttachments("%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueAttachment(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	client.httpclient = &mockHTTPClient{}

	err := client.GetIssueAttachment("SRE-1", 1, &bytes.Buffer{})
	if err != nil {
		log.Fatalf("Unexpected error: %s in test", err)
	}
}

func TestGetIssueAttachmentFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/SRE-1/attachments/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.GetIssueAttachment("SRE-1", 1, &bytes.Buffer{}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueAttachmentInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if err := client.GetIssueAttachment("%", 1, &bytes.Buffer{}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteIssueAttachment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/attachments/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONAttachment); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteIssueAttachment("BLG-1", 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestAttachment()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteIssueAttachmentFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/attachments/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteIssueAttachment("BLG-1", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteIssueAttachmentInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.DeleteIssueAttachment("%", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueParticipants(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/participants", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf(`[%s]`, testJSONUser)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetIssueParticipants("BLG-1")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*User{getTestUser()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetIssueParticipantsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/participants", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetIssueParticipants("BLG-1")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueParticipantsInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetIssueParticipants("%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueSharedFiles(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/sharedFiles", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf(`[%s]`, testJSONSharedFile)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetIssueSharedFiles("BLG-1")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*SharedFile{getTestSharedFile()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetIssueSharedFilesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/sharedFiles", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetIssueSharedFiles("BLG-1")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueSharedFilesInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetIssueSharedFiles("%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateIssueSharedFiles(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/sharedFiles", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf(`[%s]`, testJSONSharedFile)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateIssueSharedFiles("BLG-1", &CreateIssueSharedFilesInput{
		FileIDs: []int{454403},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*SharedFile{getTestSharedFile()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateIssueSharedFilesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/sharedFiles", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateIssueSharedFiles("BLG-1", &CreateIssueSharedFilesInput{
		FileIDs: []int{454403},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateIssueSharedFilesInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.CreateIssueSharedFiles("%", &CreateIssueSharedFilesInput{
		FileIDs: []int{454403},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteIssueSharedFile(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/sharedFiles/454403", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONSharedFile); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteIssueSharedFile("BLG-1", 454403)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestSharedFile()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteIssueSharedFileFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues/BLG-1/sharedFiles/454403", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteIssueSharedFile("BLG-1", 454403)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteIssueSharedFileInvalidIssueKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.DeleteIssueSharedFile("%", 454403)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
