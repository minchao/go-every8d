package every8d

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type DeliveryStatus struct {
	Count int
}

// GetDeliveryStatus retrieves the delivery status.
func (c *Client) GetDeliveryStatus(ctx context.Context, batchID, pageNo string) (*DeliveryStatus, error) {
	q := url.Values{}
	q.Set("BID", batchID)
	q.Set("PNO", pageNo)
	u, _ := url.Parse("API21/HTTP/getDeliveryStatus.ashx")

	req, err := c.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	fn := func(body io.Reader, v interface{}) error {
		// TODO
		return nil
	}

	result := new(DeliveryStatus)
	_, err = c.Do(ctx, req, fn, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
