package hltb

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
)

type mockHttpRequestProviderError struct{}

func (p *mockHttpRequestProviderError) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return nil, errors.New("test error")
}

func TestNewNilHttpClent(t *testing.T) {
	if _, err := New(nil); err == nil {
		t.Error("missing underlying client, expected error")
	}
}

func TestNewNonNilHttpClent(t *testing.T) {
	if _, err := New(&http.Client{}); err != nil {
		t.Error("has underlying client, unexpected error")
	}
}

func TestNewSearchRequestEmptySearchTerm(t *testing.T) {
	client, _ := New(&http.Client{})

	_, err := client.newSearchRequest(SearchArgs{
		Term: "  ",
	})

	if err == nil {
		t.Error("expected error")
	}
}

func TestNewSearchRequestDefaultPage(t *testing.T) {
	client, _ := New(&http.Client{})

	req, err := client.newSearchRequest(SearchArgs{
		Term: "term",
		Page: 0,
	})

	if err != nil {
		t.Error("unexpected error")
	}

	defer req.Body.Close()

	var actual searchRequest
	json.NewDecoder(req.Body).Decode(&actual)

	if actual.SearchPage != 1 {
		t.Errorf("unexpected error page number. Expected: %d, Actual: %d", 1, actual.SearchPage)
	}
}

func TestNewSearchRequestDefaultPageSize(t *testing.T) {
	client, _ := New(&http.Client{})

	req, err := client.newSearchRequest(SearchArgs{
		Term:     "term",
		PageSize: 0,
	})

	if err != nil {
		t.Error("unexpected error")
	}

	defer req.Body.Close()

	var actual searchRequest
	json.NewDecoder(req.Body).Decode(&actual)

	if actual.Size != 20 {
		t.Errorf("unexpected error page size. Expected: %d, Actual: %d", 1, actual.Size)
	}
}

func TestNewSearchRequestNewRequestError(t *testing.T) {
	client, _ := New(&http.Client{})

	client.rp = &mockHttpRequestProviderError{}

	_, err := client.newSearchRequest(SearchArgs{
		Term: "Term",
	})

	if err == nil {
		t.Error("expected error")
	}
}
