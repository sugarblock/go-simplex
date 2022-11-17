package simplex

import (
	"context"
	"errors"
	"fmt"

	v2 "github.com/sugarblock/go-simplex/api/v2"
	"github.com/sugarblock/go-simplex/types"
)

func (c *Client) GetAllEvents(ctx context.Context) (*v2.EventResponse, error) {
	queryPath := "/v2/events"
	method := "GET"

	req, err := c.newRequest(method, queryPath, nil)
	if err != nil {
		return nil, err
	}

	var eventResponse v2.EventResponse

	err = c.do(ctx, req, &eventResponse)
	if err != nil {
		if simplexErr := new(v2.SimplexError); !errors.As(err, &simplexErr) {
			var respErr types.ResponseError

			msg := fmt.Sprintf("response error: %s", err.Error())
			respErr.Message = &msg

			return nil, &respErr
		}
		return nil, err
	}

	return &eventResponse, nil
}

func (c *Client) DeleteEvent(ctx context.Context, reqv2 *v2.EventRequest) (*v2.EventDeleteResponse, error) {
	queryPath := "/v2/events"
	method := "DELETE"

	req, err := c.newRequest(method, queryPath+"/"+reqv2.EventId, nil)
	if err != nil {
		return nil, err
	}

	var eventDeleteResponse v2.EventDeleteResponse

	err = c.do(ctx, req, &eventDeleteResponse)
	if err != nil {
		if simplexErr := new(v2.SimplexError); !errors.As(err, &simplexErr) {
			var respErr types.ResponseError

			msg := fmt.Sprintf("response error: %s", err.Error())
			respErr.Message = &msg

			return nil, &respErr
		}
		return nil, err
	}

	return &eventDeleteResponse, nil
}
