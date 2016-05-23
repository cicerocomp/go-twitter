package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// SearchMetadata ...
type SearchMetadata struct {
	CompletedIn float64 `json:"completed_in"`
	Count       int     `json:"count"`
	MaxID       int     `json:"max_id"`
	MaxIDStr    string  `json:"max_id_str"`
	NextResults string  `json:"next_results"`
	Query       string  `json:"query"`
	RefreshURL  string  `json:"refresh_url"`
	SinceID     int     `json:"since_id"`
	SinceIDStr  string  `json:"since_id_str"`
}

// Search ...
type Search struct {
	SearchMetadata *SearchMetadata `json:"search_metadata"`
	Statuses       []Tweet         `json:"statuses"`
}

// SearchService provides a method for search in old tweets
type SearchService struct {
	sling *sling.Sling
}

// newSearchService returns a new SearchService.
func newSearchService(sling *sling.Sling) *SearchService {
	return &SearchService{
		sling: sling.Path("search/"),
	}
}

// SearchTweetsParams are the params for SearchService.Tweets.
type SearchTweetsParams struct {
	Q          string `url:"q,omitempty"`
	GeoCode    string `url:"geocode,omitempty"`
	Lang       string `url:"lang,omitempty"`
	ResultType string `url:"result_type,omitempty"`
	Count      *int64 `url:"count,omitempty"`
	Locale     string `url:"locale,omitempty"`
	MaxID      *int64 `url:"max_id,omitempty"`
	SinceID    *int64 `url:"since_id,omitempty"`
}

// Tweets Returns a collection of relevant Tweets matching a specified query.
// https://dev.twitter.com/rest/reference/get/search/tweets
func (s *SearchService) Tweets(params *SearchTweetsParams) (*Search, *http.Response, error) {
	search := new(Search)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("tweets.json").QueryStruct(params).Receive(search, apiError)
	return search, resp, relevantError(err, *apiError)
}
