package backlog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func getTestIssues() []Issue {
	return []Issue{
		getTestIssuesWithID(1),
		getTestIssuesWithID(2),
		getTestIssuesWithID(3),
		getTestIssuesWithID(4),
	}
}

func getTestIssuesWithID(id int) Issue {
	return Issue{
		ID:        id,
		ProjectID: 1,
		IssueKey:  "BLG-1",
		KeyID:     1,
		IssueType: IssueType{
			ID:           2,
			ProjectID:    1,
			Name:         "タスク",
			Color:        "#7ea800",
			DisplayOrder: 0,
		},
		Summary:     "first issue",
		Description: "",
		Resolutions: "",
		Priority: Priority{
			ID:   3,
			Name: "中",
		},
		Status: Status{
			ID:           1,
			ProjectID:    1,
			Name:         "未対応",
			Color:        "#ed8077",
			DisplayOrder: 1000,
		},
		Assignee: User{
			ID:          2,
			Name:        "eguchi",
			RoleType:    2,
			Lang:        "",
			MailAddress: "eguchi@nulab.example",
		},
		// Category: []int{},
		// Versions: []int{},
		Milestone: []Milestone{
			{
				ID:        30,
				ProjectID: 1,
			},
		},
		StartDate:      "",
		DueDate:        "",
		EstimatedHours: "",
		ActualHours:    "",
		ParentIssueID:  nil,
		CreatedUser: User{
			ID:          1,
			UserID:      "admin",
			Name:        "admin",
			RoleType:    1,
			Lang:        "ja",
			MailAddress: "eguchi@nulab.example",
		},
		Created: JSONTime("2012-07-23T06:10:15Z"),
		UpdatedUser: User{
			ID:          1,
			UserID:      "admin",
			Name:        "admin",
			RoleType:    1,
			Lang:        "ja",
			MailAddress: "eguchi@nulab.example",
		},
		Updated:      JSONTime("2013-02-07T08:09:49Z"),
		CustomFields: nil,
		Attachments: []Attachment{
			{
				ID:   1,
				Name: "IMGP0088.JPG",
				Size: 85079,
			},
		},
		SharedFiles: []SharedFile{
			{
				ID:   454403,
				Type: "file",
				Dir:  "/ユーザアイコン/",
				Name: "01_サラリーマン.png",
				Size: 2735,
				CreatedUser: User{
					ID:          5686,
					UserID:      "takada",
					Name:        "takada",
					RoleType:    2,
					Lang:        "ja",
					MailAddress: "takada@nulab.example",
				},
				Created: JSONTime("2009-02-27T03:26:15Z"),
				UpdatedUser: User{
					ID:          5686,
					UserID:      "takada",
					Name:        "takada",
					RoleType:    2,
					Lang:        "ja",
					MailAddress: "takada@nulab.example",
				},
				Updated: JSONTime("2009-03-03T16:57:47Z"),
			},
		},
		Stars: []Star{
			{
				ID:      10,
				Comment: "",
				URL:     "https://xx.backlog.jp/view/BLG-1",
				Title:   "[BLG-1] first issue | 課題の表示 - Backlog",
				Presenter: User{
					ID:          2,
					UserID:      "eguchi",
					Name:        "eguchi",
					RoleType:    2,
					Lang:        "ja",
					MailAddress: "eguchi@nulab.example",
				},
				Created: JSONTime("2013-07-08T10:24:28Z"),
			},
		},
	}
}

func getIssues(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestIssues())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func TestGetIssues(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/issues", getIssues)
	expected := getTestIssues()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetIssuesInput{
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
		ParentChild:    11,
		Sort:           SortIssueType,
		Offset:         10,
		CreatedSince:   "2019-01-07",
		CreatedUntil:   "2020-01-07",
		UpdatedSince:   "2019-01-07",
		UpdatedUntil:   "2020-01-07",
		StartDateSince: "2019-01-07",
		StartDateUntil: "2020-01-07",
		DueDateSince:   "2019-01-07",
		DueDateUntil:   "2020-01-07",
		IDs:            []int{1},
		ParentIssueIDs: []int{11, 12, 13},
		Keyword:        "test",
	}
	issues, err := api.GetIssues(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, issues) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetIssuesWithOrderAndCount(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/issues", getIssues)
	expected := getTestIssues()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetIssuesInput{
		Order: OrderAsc,
		Count: 100,
		Sort:  SortCategory,
	}
	issues, err := api.GetIssues(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, issues) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetIssuesFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/issues", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetIssuesInput{
		Sort: SortVersion,
	}
	_, err := api.GetIssues(input)
	if err == nil {
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
