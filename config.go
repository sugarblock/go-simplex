package simplex

import (
	"fmt"
	"net/url"
	"os"

	"github.com/sugarblock/go-simplex/types"
)

const (
	defaultBaseURL          = "https://sandbox.test-simplexcc.com/wallet/merchant"
	defaultAuthHeaderPrefix = "apiKey"
)

type Config struct {
	URL              *url.URL
	HeaderAuthPrefix string
	ApiKey           string
}

func newConfigFromEnv() (*Config, error) {
	simplexUrl := os.Getenv("SIMPLEX_URL")
	if simplexUrl == "" {
		simplexUrl = defaultBaseURL
	}

	url, err := ParseBaseURL(simplexUrl)
	if err != nil {
		return nil, err
	}

	headerAuthPrefix := os.Getenv("SIMPLEX_AUTHORIZATION_HEADER_PREFIX")
	if headerAuthPrefix == "" {
		headerAuthPrefix = defaultAuthHeaderPrefix
	}

	apiKey := os.Getenv("SIMPLEX_APIKEY")
	if apiKey == "" {
		msg := "you must provide an apiKey"
		return nil, &types.EnvError{Message: &msg}
	}

	return &Config{
		URL:              url,
		HeaderAuthPrefix: headerAuthPrefix,
		ApiKey:           apiKey,
	}, nil
}

func ParseBaseURL(basURL string) (*url.URL, error) {
	url, err := url.Parse(basURL)
	if err != nil {
		msg := fmt.Sprintf("parsing URL: %s", err.Error())
		return nil, &types.ParsingUrlError{Message: &msg}
	}
	return url, nil
}
