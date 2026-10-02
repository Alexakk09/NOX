package websearch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SearXNGSearcher struct {
	client  *http.Client
	baseURL string
}

func NewSearXNGSearcher(baseURL string) *SearXNGSearcher {
	return &SearXNGSearcher{
		client:  &http.Client{},
		baseURL: baseURL,
	}
}

type searXNGResponse struct {
	Results []struct {
		Title   string `json:"title"`
		URL     string `json:"url"`
		Content string `json:"content"`
	} `json:"results"`
}

func (s *SearXNGSearcher) Search(query string) ([]Result, error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("format", "json")

	searchURL := s.baseURL + "/search?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"searxng search failed with status %d",
			response.StatusCode,
		)
	}

	var data searXNGResponse

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	results := make([]Result, 0, len(data.Results))

	for _, result := range data.Results {
		results = append(results, Result{
			Title:   result.Title,
			URL:     result.URL,
			Snippet: result.Content,
		})
	}

	return results, nil
}