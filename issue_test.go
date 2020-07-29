package backlog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/pkg/errors"

	"github.com/kylelemons/godebug/pretty"
)

const testJSONIssue string = `{
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
	"resolutions": "",
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
	"assignee": {
		"id": 2,
		"name": "eguchi",
		"roleType": 2,
		"lang": null,
		"mailAddress": "eguchi@nulab.example"
	},
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
	"updated": "2006-01-02T15:04:05Z",
	"customFields": [],
	"attachments": [
		{
			"id": 1,
			"name": "IMGP0088.JPG",
			"size": 85079
		}
	],
	"sharedFiles": [
		{
			"id": 454403,
			"type": "file",
			"dir": "/ユーザアイコン/",
			"name": "01_サラリーマン.png",
			"size": 2735,
			"createdUser": {
				"id": 5686,
				"userId": "takada",
				"name": "takada",
				"roleType": 2,
				"lang": "ja",
				"mailAddress": "takada@nulab.example"
			},
			"created": "2006-01-02T15:04:05Z",
			"updatedUser": {
				"id": 5686,
				"userId": "takada",
				"name": "takada",
				"roleType": 2,
				"lang": "ja",
				"mailAddress": "takada@nulab.example"
			},
			"updated": "2006-01-02T15:04:05Z"
		}
	],
	"stars": [
		{
			"id": 10,
			"comment": null,
			"url": "https://xx.backlog.jp/view/BLG-1",
			"title": "[BLG-1] first issue | 課題の表示 - Backlog",
			"presenter": {
				"id": 2,
				"userId": "eguchi",
				"name": "eguchi",
				"roleType": 2,
				"lang": "ja",
				"mailAddress": "eguchi@nulab.example"
			},
			"created": "2006-01-02T15:04:05Z"
		}
	]
}`

func getTestIssuesWithID(id int) *Issue {
	return &Issue{
		ID:        Int(id),
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
		Resolutions: String(""),
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
		Assignee: &User{
			ID:          Int(2),
			Name:        String("eguchi"),
			RoleType:    RoleType(2),
			Lang:        nil,
			MailAddress: String("eguchi@nulab.example"),
		},
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
		CreatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    1,
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
		Updated:      &Timestamp{referenceTime},
		CustomFields: []*CustomField{},
		Attachments: []*Attachment{
			{
				ID:   Int(1),
				Name: String("IMGP0088.JPG"),
				Size: Int(85079),
			},
		},
		SharedFiles: []*SharedFile{
			{
				ID:   Int(454403),
				Type: String("file"),
				Dir:  String("/ユーザアイコン/"),
				Name: String("01_サラリーマン.png"),
				Size: Int(2735),
				CreatedUser: &User{
					ID:          Int(5686),
					UserID:      String("takada"),
					Name:        String("takada"),
					RoleType:    2,
					Lang:        String("ja"),
					MailAddress: String("takada@nulab.example"),
				},
				Created: &Timestamp{referenceTime},
				UpdatedUser: &User{
					ID:          Int(5686),
					UserID:      String("takada"),
					Name:        String("takada"),
					RoleType:    2,
					Lang:        String("ja"),
					MailAddress: String("takada@nulab.example"),
				},
				Updated: &Timestamp{referenceTime},
			},
		},
		Stars: []*Star{
			{
				ID:      Int(10),
				Comment: nil,
				URL:     String("https://xx.backlog.jp/view/BLG-1"),
				Title:   String("[BLG-1] first issue | 課題の表示 - Backlog"),
				Presenter: &User{
					ID:          Int(2),
					UserID:      String("eguchi"),
					Name:        String("eguchi"),
					RoleType:    2,
					Lang:        String("ja"),
					MailAddress: String("eguchi@nulab.example"),
				},
				Created: &Timestamp{referenceTime},
			},
		},
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

	want := []*Issue{getTestIssuesWithID(1)}
	if !reflect.DeepEqual(want, issues) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetIssuesWithOrderAndCount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONIssue)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	opts := &GetIssuesOptions{
		Order: OrderAsc,
		Count: Int(100),
		Sort:  SortCategory,
	}

	issues, err := client.GetIssues(opts)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Issue{getTestIssuesWithID(1)}
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
		{getTestIssuesWithID(1)},
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
