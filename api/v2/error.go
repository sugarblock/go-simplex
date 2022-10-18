package v2

import "fmt"

type SimplexError struct {
	StatusCode   *int                  `json:"statusCode,omitempty"`
	ErrorMessage string                `json:"error,omitempty"`
	Errors       *[]SimplexParamsError `json:"errors,omitempty"`
}

type SimplexParamsError struct {
	Code    *string   `json:"code,omitempty"`
	Params  *[]string `json:"params,omitempty"`
	Message *string   `json:"message,omitempty"`
	Path    *string   `json:"path,omitempty"`
}

func (e SimplexError) Error() string {
	return fmt.Sprintf("StatusCode: %d:\n Error:%s\n Errors: %#v", e.StatusCode, e.ErrorMessage, e.Errors)
}
