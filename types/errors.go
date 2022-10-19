package types

import (
	"encoding/json"
	"fmt"
)

type ResponseError struct {
	Message *string
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}

func (e *ResponseError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}

func (e *ResponseError) ErrorCode() string { return "ResponseError" }

type RequestError struct {
	Message *string
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}

func (e *RequestError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}

func (e *RequestError) ErrorCode() string { return "RequestError" }

type ParsingUrlError struct {
	Message *string
}

func (e *ParsingUrlError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}

func (e *ParsingUrlError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}

func (e *ParsingUrlError) ErrorCode() string { return "ParsingUrlError" }

type EnvError struct {
	Message *string
}

func (e *EnvError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}

func (e *EnvError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}

func (e *EnvError) ErrorCode() string { return "EnvError" }

type ValidationError struct {
	Messages *[]string `json:"messages,omitempty"`
}

func (e ValidationError) Error() string {
	out, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s: %s", e.ErrorCode(), string(out))
}

func (e *ValidationError) ErrorCode() string { return "ValidationError" }
