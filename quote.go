package simplex

import (
	"context"

	v2 "github.com/sugarblock/go-simplex/api/v2"
)

func (c *Client) GetQuote(ctx context.Context, reqv2 *v2.QuoteRequest) (*v2.QuoteResponse, error) {
	queryPath := "/v2/quote"
	method := "POST"

	req, err := c.newRequest(method, queryPath, reqv2)
	if err != nil {
		return nil, err
	}

	var quoteResponse v2.QuoteResponse
	err = c.do(ctx, req, &quoteResponse)
	if err != nil {
		return nil, err
	}

	return &quoteResponse, nil
}
