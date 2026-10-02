package websearch

type Result struct {
	Title   string
	URL     string
	Snippet string
}

type Searcher interface {
	Search(query string) ([]Result, error)
}