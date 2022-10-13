package types

import "fmt"

type ErrorResponse struct {
	StatusCode int      `json:"statusCode"`
	Messages   []string `json:"messages"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("StatusCode: %d:\n %+q", e.StatusCode, e.Messages)
}
