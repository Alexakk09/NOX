package clients



type APIClient struct {
	BaseURL string
	Token   string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
	}
}

