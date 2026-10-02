package websearch

import "testing"

func TestSearXNGSearch(t *testing.T) {
	searcher := NewSearXNGSearcher("http://localhost:8888")

	results, err := searcher.Search("Go programming language")
	if err != nil {
		t.Fatal(err)
	}

	if len(results) == 0 {
		t.Fatal("expected at least one search result")
	}

	for _, result := range results {
		t.Logf("Title: %s", result.Title)
		t.Logf("URL: %s", result.URL)
		t.Logf("Snippet: %s", result.Snippet)
	}
}
