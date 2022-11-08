package ewhs

import "fmt"

// BaseError contains the general error structure
// returned by mollie.
type BaseError struct {
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// Error interface compliance.
func (be *BaseError) Error() string {
	//str := fmt.Sprintf("%d - %s: %s", be.Status, be.Title, be.Detail)
	str := fmt.Sprintf("%d - %s", be.Status, be.Title)

	return str
}
