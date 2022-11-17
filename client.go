package simplex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	v2 "github.com/sugarblock/go-simplex/api/v2"
)

const (
	defaultTimeout = 60 * time.Second
)

type Client struct {
	client  *http.Client
	rootURL string
	bearer  string
}

func NewEnvClient() (*Client, error) {
	config, err := newConfigFromEnv()
	if err != nil {
		return nil, err
	}

	baseURL := config.URL.String()
	authHeaderPrefix := config.HeaderAuthPrefix
	apiKey := config.ApiKey

	return NewClient(nil, baseURL, authHeaderPrefix, apiKey)
}

func NewClient(client *http.Client, baseURL, authHeaderPrefix, apiKey string) (*Client, error) {
	if client == nil {
		transport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: defaultTimeout,
			}).DialContext,
			TLSHandshakeTimeout: defaultTimeout,
		}
		client = &http.Client{
			Timeout:   defaultTimeout,
			Transport: transport,
		}
	}

	simplex := new(Client)
	simplex.client = client

	url, err := ParseBaseURL(baseURL)
	if err != nil {
		return nil, err
	}

	if apiKey == "" {
		return nil, fmt.Errorf("apiKey must not be empty")
	}

	if authHeaderPrefix == "" {
		return nil, fmt.Errorf("authHeaderPrefix must not be empty")
	}

	simplex.rootURL = url.String()
	simplex.bearer = authHeaderPrefix + " " + apiKey

	return simplex, nil
}

func (c *Client) newRequest(method, resource string, body interface{}) (*http.Request, error) {
	if resource == "" {
		return nil, fmt.Errorf("resource can't be nil")
	}

	var b []byte
	if body != nil {
		var err error
		b, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	reqURL := c.rootURL + resource

	req, err := http.NewRequest(method, reqURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.bearer)
	req.Header.Set("Accept", "application/json")

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

	defer resp.Body.Close()

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(body) == 0 {
		return fmt.Errorf("check if server supports the requested URL: %s", reqURL)
	}

	var simplexErr v2.SimplexError
	err = json.Unmarshal(body, &simplexErr)
	if err != nil {
		return err
	}

	simplexErr.StatusCode = &resp.StatusCode

	return &simplexErr
}
