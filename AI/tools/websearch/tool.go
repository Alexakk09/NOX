package websearch

import (
	"Jarvis/tools/registry"
	"errors"
)

var ErrInvalidQuery = errors.New("query is required")

type Tool struct {
	searcher Searcher
}

func NewTool(searcher Searcher) *Tool {
	return &Tool{
		searcher: searcher,
	}
}

func (t *Tool) Name() string {
	return "web_search"
}

func (t *Tool) Description() string {
	return "Search the web for current information."
}

func (t *Tool) Execute(arguments map[string]any) (any, error) {
	query, ok := arguments["query"].(string)
	if !ok {
		return nil, ErrInvalidQuery
	}

	return t.searcher.Search(query)
}

var _ registry.Tool = (*Tool)(nil)

func (t *Tool) Parameters() []registry.Parameter {
	return []registry.Parameter{
		{
			Name:        "query",
			Type:        "string",
			Description: "The search query.",
			Required:    true,
		},
	}
}
