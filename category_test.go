package backlog

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

const testJSONCategory string = `{
	"id": 12,
	"name": "開発",
	"displayOrder": 0
}`

func getTestCategory() *Category {
	return &Category{
		ID:           Int(12),
		Name:         String("開発"),
		DisplayOrder: Int(0),
	}
}

func TestGetCategories(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/categories", func(w http.ResponseWriter, _ *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONCategory)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	categories, err := client.GetCategories("SRE")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Category{getTestCategory()}
	if !reflect.DeepEqual(want, categories) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, categories)))
	}
}

func TestGetCategoriesInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetCategories("%%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateCategoriesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/categories", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetCategories("SRE")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateCategory(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/categories", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONCategory); err != nil {
			t.Fatal(err)
		}
	})

	input := &CreateCategoryInput{
		Name: String("開発"),
	}
	category, err := client.CreateCategory("SRE", input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestCategory()
	if !reflect.DeepEqual(want, category) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, category)))
	}
}

func TestCreateCategoryFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/categories", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &CreateCategoryInput{
		Name: String("開発"),
	}
	_, err := client.CreateCategory("SRE", input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateCategoryInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.CreateCategory("%%", &CreateCategoryInput{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateCategory(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/categories/1", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONCategory); err != nil {
			t.Fatal(err)
		}
	})

	input := &UpdateCategoryInput{
		Name: String("開発"),
	}
	category, err := client.UpdateCategory("SRE", 1, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestCategory()
	if !reflect.DeepEqual(want, category) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, category)))
	}
}

func TestUpdateCategoryFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/categories/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &UpdateCategoryInput{
		Name: String("開発"),
	}
	_, err := client.UpdateCategory("SRE", 1, input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateCategoryInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.UpdateCategory("%%", 1, &UpdateCategoryInput{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteCategory(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/categories/1", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONCategory); err != nil {
			t.Fatal(err)
		}
	})

	category, err := client.DeleteCategory("SRE", 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestCategory()
	if !reflect.DeepEqual(want, category) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, category)))
	}
}

func TestDeleteCategoryFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/categories/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteCategory("SRE", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteCategoryInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.DeleteCategory("%%", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetCategories_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.GetCategories("SRE")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestCreateCategory_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.CreateCategory("SRE", &CreateCategoryInput{Name: String("test")})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateCategory_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.UpdateCategory("SRE", 1, &UpdateCategoryInput{Name: String("test")})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteCategory_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.DeleteCategory("SRE", 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
