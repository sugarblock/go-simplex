package simplex

import (
	"context"

	v2 "github.com/sugarblock/go-simplex/api/v2"
)

func (c *Client) GetPayment(ctx context.Context, reqv2 *v2.PaymentRequest) (*v2.PaymentResponse, error) {
	queryPath := "/v2/payments/partner/data"
	method := "POST"

	req, err := c.newRequest(method, queryPath, reqv2)
	if err != nil {
		return nil, err
	}

	var paymentResponse v2.PaymentResponse
	err = c.do(ctx, req, &paymentResponse)
	if err != nil {
		return nil, err
	}

	return &paymentResponse, nil
}
