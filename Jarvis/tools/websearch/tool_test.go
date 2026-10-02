package websearch

import (
	"testing"
)

func TestWebSearchTool(t *testing.T) {
	searcher := NewSearXNGSearcher("http://localhost:8888")
	tool := NewTool(searcher)

	result, err := tool.Execute(map[string]any{
		"query": "Go programming language",
	})
	if err != nil {
		t.Fatal(err)
	}

	results, ok := result.([]Result)
	if !ok {
		t.Fatalf("expected []Result, got %T", result)
	}

	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}

	t.Logf("Found %d results", len(results))
	t.Logf("First result: %+v", results[0])
}