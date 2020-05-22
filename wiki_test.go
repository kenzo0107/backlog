package backlog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func getTestWiki() Wiki {
	return getTestWikiWithID(1)
}

func getTestWikis() []Wiki {
	return []Wiki{
		getTestWikiWithID(1),
		getTestWikiWithID(2),
		getTestWikiWithID(3),
		getTestWikiWithID(4),
	}
}

func getTestWikiCount() Page {
	return Page{
		Count: len(getTestWikis()),
	}
}

func getTestWikiWithID(id int) Wiki {
	return Wiki{
		ID:        id,
		ProjectID: 1,
		Name:      "Home",
		Content:   "test",
		Tags: []Tag{
			{
				ID:   12,
				Name: "議事録",
			},
		},
		Attachments: []Attachment{
			{
				ID:   1,
				Name: "test.json",
				Size: 8857,
				CreatedUser: User{
					ID:          1,
					UserID:      "admin",
					RoleType:    1,
					Lang:        "ja",
					MailAddress: "eguchi@nulab.example",
				},
				Created: JSONTime("2020-02-19T05:54:32Z"),
			},
		},
	}
}

func getTestWikiTagWithID(id int) Tag {
	return Tag{
		ID:   id,
		Name: "test",
	}
}

func getTestWikiTags() []Tag {
	return []Tag{
		getTestWikiTagWithID(1),
		getTestWikiTagWithID(2),
	}
}

func getWikis(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestWikis())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func getWikiCount(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestWikiCount())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func getWikiTags(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestWikiTags())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func getWiki(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestWiki())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func TestGetWikis(t *testing.T) {
	http.HandleFunc("/api/v2/wikis", getWikis)
	expected := getTestWikis()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	wikis, err := api.GetWikis(1, "test")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, wikis) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetWikisWithInvalidProjectKey(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetWikis(true, "test")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikisFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/wikis", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	if _, err := api.GetWikis(1, "test"); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiCount(t *testing.T) {
	http.HandleFunc("/api/v2/wikis/count", getWikiCount)
	expected := getTestWikiCount()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	count, err := api.GetWikiCount(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected.Count, count) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetWikiCountFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/wikis/count", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetWikiCount(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiCountWithInvalidProjectKey(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetWikiCount(true)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiTags(t *testing.T) {
	http.HandleFunc("/api/v2/wikis/tags", getWikiTags)
	expected := getTestWikiTags()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	tags, err := api.GetWikiTags(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, tags) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetWikiTagsWithInvalidProjectKey(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetWikiTags(true)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiTagsFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/wikis/tags", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetWikiTags(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWikiByID(t *testing.T) {
	http.HandleFunc("/api/v2/wikis/1", getWiki)
	expected := getTestWiki()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	wiki, err := api.GetWiki(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, wiki) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetWikiByIDFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/wikis/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetWiki(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateWiki(t *testing.T) {
	http.HandleFunc("/api/v2/wikis", getWiki)
	expected := getTestWiki()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &CreateWikiInput{
		ProjectID: 1,
		Name:      "Home",
		Content:   "test",
	}
	wiki, err := api.CreateWiki(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, wiki) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateWikiFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/wikis", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &CreateWikiInput{
		ProjectID: 1,
		Name:      "Home",
		Content:   "test",
	}
	_, err := api.CreateWiki(input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateWiki(t *testing.T) {
	http.HandleFunc("/api/v2/wikis/1", getWiki)
	expected := getTestWiki()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &UpdateWikiInput{
		WikiID:  1,
		Name:    "Home",
		Content: "test",
	}
	wiki, err := api.UpdateWiki(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, wiki) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateWikiFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/wikis/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &UpdateWikiInput{
		WikiID:  1,
		Name:    "Home",
		Content: "test",
	}
	_, err := api.UpdateWiki(input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteWiki(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/wikis/1", getWiki)
	expected := getTestWiki()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	wiki, err := api.DeleteWiki(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, wiki) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteWikiFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/wikis/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.DeleteWiki(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
