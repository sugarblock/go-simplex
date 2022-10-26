package simplex

import (
	"context"
	"errors"
	"fmt"

	v2 "github.com/sugarblock/go-simplex/api/v2"
	"github.com/sugarblock/go-simplex/types"
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
		if simplexErr := new(v2.SimplexError); !errors.As(err, &simplexErr) {
			var respErr types.ResponseError

			msg := fmt.Sprintf("response error: %s", err.Error())
			respErr.Message = &msg

			return nil, &respErr
		}
		return nil, err
	}

	return &quoteResponse, nil
}
