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
	v2 "github.com/sugarblock/go-simplex/api/v2"
	"github.com/sugarblock/go-simplex/types"
)

const (
	defaultBaseURL          = "https://sandbox.test-simplexcc.com/wallet/merchant"
	defaultAuthHeaderPrefix = "apiKey"
	defaultTimeout          = 60 * time.Second
)

type Client struct {
	client  *http.Client
	rootURL string
	apiKey  string
}

var validate *validator.Validate

func NewClient(client *http.Client, baseURL, authHeaderPrefix, apiKey *string) (*Client, error) {

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

	rootURL, err := getValue(baseURL, defaultBaseURL, "SIMPLEX_URL")
	if err != nil {
		return nil, fmt.Errorf("reading simplex url: %w", err)
	}

	url, err := url.ParseRequestURI(rootURL)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}

	simplex.rootURL = url.String()

	authPrefixHeaderValue, err := getValue(authHeaderPrefix, defaultAuthHeaderPrefix, "SIMPLEX_AUTHORIZATION_HEADER_PREFIX")
	if err != nil {
		return nil, fmt.Errorf("reading authPrefixHeader: %w", err)
	}

	apiKeyValue, err := getValue(apiKey, "", "SIMPLEX_APIKEY")
	if err != nil {
		return nil, fmt.Errorf("reading apiKey: %w", err)
	}

	simplex.apiKey = authPrefixHeaderValue + " " + apiKeyValue

	if validate == nil {
		validate = validator.New()
	}

	return simplex, nil
}

func getValue(value *string, defaultValue, envKey string) (string, error) {
	var v string
	if value != nil {
		v = *value
	} else if vFromEnv := os.Getenv(envKey); vFromEnv != "" {
		v = vFromEnv
	} else {
		v = defaultValue
	}

	if v == "" {
		return "", fmt.Errorf("empty value not allowed")
	}
	return v, nil
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

	reqURL := c.rootURL + resource

	req, err := http.NewRequest(method, reqURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.apiKey)

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

	var respErr types.ResponseError

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		respErr.Messages = []string{err.Error()}
		return &respErr
	}

	if len(body) == 0 {
		respErr.Messages = []string{fmt.Sprintf("check if server supports the requested URL: %s", reqURL)}
		return &respErr
	}

	var simplexErr v2.SimplexError
	err = json.NewDecoder(resp.Body).Decode(simplexErr)
	if err != nil {
		return err
	}
	simplexErr.StatusCode = &resp.StatusCode

	return &simplexErr
}
