package swapi

import "net/http"

const (
	API_URL             = "https://swapi.dev/api"
	SEARCH_QUERY_PARAMS = "?search="
)

// Client a estrtura que representa 
type Client struct {
	url        string       // url is the base url for SWAPI
	httpClient *http.Client // HTTP client used to communicate with the SWAPI
}

// NewClient returns a new SWAPI client.
func NewClient() *Client {
	return &Client{
		url:        API_URL,
		httpClient: &http.Client{},
	}
}
