package models

// PageOptions represents the pagination parameters for API requests
type PageOptions struct {
	// Page is the page number of results to retrieve (1-based)
	Page int `url:"page,omitempty"`

	// PageLen is the number of results to include per page
	PageLen int `url:"pagelen,omitempty"`

	// Q is the query string to filter the results
	Q string `url:"q,omitempty"`
}
