package simplex

import (
	"context"
	"errors"
	"fmt"

	v2 "github.com/sugarblock/go-simplex/api/v2"
	"github.com/sugarblock/go-simplex/types"
)

func (c *Client) GetPayment(ctx context.Context, reqv2 *v2.PaymentRequest) (*v2.PaymentResponse, error) {
	queryPath := "/v2/payments/partner/data"
	method := "POST"

	req, err := c.newRequest(method, queryPath, reqv2)
	if err != nil {
		if validationErr := new(types.ValidationError); !errors.As(err, &validationErr) {
			var reqErr types.RequestError

			msg := fmt.Sprintf("request error: %s", err.Error())
			reqErr.Message = &msg

			return nil, &reqErr
		}
		return nil, err
	}

	var paymentResponse v2.PaymentResponse

	err = c.do(ctx, req, &paymentResponse)
	if err != nil {
		if simplexErr := new(v2.SimplexError); !errors.As(err, &simplexErr) {
			var respErr types.ResponseError

			msg := fmt.Sprintf("response error: %s", err.Error())
			respErr.Message = &msg

			return nil, &respErr
		}
		return nil, err
	}

	return &paymentResponse, nil
}
