package hltb

import "testing"

func TestGetSearchRequestPayload(t *testing.T) {

	expected := SearchArgs{
		Term:     "Marvel Vs. Capcom 2",
		Page:     2,
		PageSize: 10,
	}

	req := getSearchRequestPayload(expected)

	if len(req.SearchTerms) != 4 {
		t.Errorf("unexpected search term count. Expected: %d, Actual: %d", 4, len(req.SearchTerms))
	}

	if req.SearchPage != expected.Page {
		t.Errorf("unexpected search page number. Expected: %d, Actual: %d", expected.Page, req.SearchPage)
	}

	if req.Size != expected.PageSize {
		t.Errorf("unexpected search page size. Expected: %d, Actual: %d", expected.PageSize, req.Size)
	}
}
