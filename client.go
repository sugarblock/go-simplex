package simplex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-playground/validator"
	"github.com/sugarblock/go-simplex/types"
)

const (
	defaultBaseURL = "https://sandbox.test-simplexcc.com/wallet/merchant"
	DefaultTimeout = 60 * time.Second
)

type Client struct {
	client            *http.Client
	defaultRootURL    string
	customHTTPHeaders map[string]string
}

var validate *validator.Validate

func NewClient(baseURL *string, client *http.Client, httpHeaders map[string]string) (*Client, error) {
	if client == nil {
		transport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: DefaultTimeout,
			}).DialContext,
			TLSHandshakeTimeout: DefaultTimeout,
		}
		client = &http.Client{
			Timeout:   DefaultTimeout,
			Transport: transport,
		}
	}

	if validate == nil {
		validate = validator.New()
	}

	simplex := new(Client)
	simplex.client = client

	var rootURL string
	if baseURL != nil {
		rootURL = *baseURL
	} else if urlFromEnv := os.Getenv("SIMPLEX_URL"); urlFromEnv != "" {
		rootURL = urlFromEnv
	} else {
		rootURL = defaultBaseURL
	}

	url, err := url.ParseRequestURI(rootURL)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}

	simplex.defaultRootURL = url.String()
	simplex.customHTTPHeaders = httpHeaders

	return simplex, nil
}

func (c *Client) newRequest(method, resource string, body interface{}) (*http.Request, error) {
	if resource == "" {
		return nil, fmt.Errorf("resource can't be nil")
	}

	var b []byte
	if body != nil {
		err := validate.Struct(body)
		if err != nil {
			var errResp types.ErrorResponse
			errResp.StatusCode = -1
			messages := []string{}
			for _, err := range err.(validator.ValidationErrors) {
				messages = append(messages, fmt.Sprintf("%s:%s:%s", err.Kind().String(), err.Namespace(), err.ActualTag()))
			}
			errResp.Messages = messages
			return nil, &errResp
		}

		b, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	reqURL := c.defaultRootURL + resource

	req, err := http.NewRequest(method, reqURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	for k, v := range c.customHTTPHeaders {
		req.Header.Set(k, v)
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) error {
	var err error
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("making HTTP request: %w", err)
	}

	err = c.checkResponse(resp, req.URL.RawPath)
	if err != nil {
		return err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return err
		}
	}

	return err
}

func (c *Client) checkResponse(resp *http.Response, reqURL string) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return nil
	}

	var errResp types.ErrorResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errResp.StatusCode = -1
		errResp.Messages = []string{err.Error()}
		return &errResp
	}

	if len(body) == 0 {
		errResp.StatusCode = http.StatusBadRequest
		errResp.Messages = []string{fmt.Sprintf("check if server supports the requested  URL: %s", reqURL)}
		return &errResp
	}

	errResp.StatusCode = resp.StatusCode
	errResp.Messages = []string{resp.Status}

	return &errResp
}
