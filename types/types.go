package types

import (
	"fmt"

	v2 "github.com/sugarblock/go-simplex/api/v2"
)

type ResponseError struct {
	Messages     []string         `json:"messages,omitempty"`
	SimplexError *v2.SimplexError `json:"simplexError,omitempty"`
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("StatusCode: %#v:\n %+q", e.SimplexError, e.Messages)
}
