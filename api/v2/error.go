package v2

import (
	"encoding/json"
	"fmt"
)

type SimplexError struct {
	StatusCode   *int                  `json:"statusCode,omitempty"`
	ErrorMessage *string               `json:"error,omitempty"`
	Errors       *[]SimplexParamsError `json:"errors,omitempty"`
}

type SimplexParamsError struct {
	Code    *string   `json:"code,omitempty"`
	Params  *[]string `json:"params,omitempty"`
	Message *string   `json:"message,omitempty"`
	Path    *string   `json:"path,omitempty"`
}

func (e SimplexError) Error() string {
	out, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s: %s", e.ErrorCode(), string(out))
}

func (e *SimplexError) ErrorCode() string { return "SimplexError" }
