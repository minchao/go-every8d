package every8d

import (
	"context"
	"encoding/csv"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// DeliveryStatus represents delivery status.
type DeliveryStatus struct {
	Name     string
	Mobile   string
	SendTime string
	Cost     float64
	Status   StatusCode
}

// DeliveryStatusResponse represents the response of get delivery status.
type DeliveryStatusResponse struct {
	Count   int
	Records []DeliveryStatus
}

// GetDeliveryStatus retrieves the delivery status.
func (c *Client) GetDeliveryStatus(ctx context.Context, batchID, pageNo string) (*DeliveryStatusResponse, error) {
	return c.getDeliveryStatus(ctx, "API21/HTTP/getDeliveryStatus.ashx", batchID, pageNo)
}

// GetMMSDeliveryStatus retrieves the MMS delivery status.
func (c *Client) GetMMSDeliveryStatus(ctx context.Context, batchID, pageNo string) (*DeliveryStatusResponse, error) {
	return c.getDeliveryStatus(ctx, "API21/HTTP/MMS/getDeliveryStatus.ashx", batchID, pageNo)
}

func (c *Client) getDeliveryStatus(ctx context.Context, urlStr, batchID, pageNo string) (*DeliveryStatusResponse, error) {
	q := url.Values{}
	q.Set("BID", batchID)
	q.Set("PNO", pageNo)
	u, _ := url.Parse(urlStr)
	u.RawQuery = q.Encode()

	req, err := c.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	fn := func(body io.Reader, v interface{}) error {
		response := DeliveryStatusResponse{}

		r := csv.NewReader(body)
		r.Comma = '\t'

		if record, err := r.Read(); err == nil && len(record) == 1 {
			response.Count, _ = strconv.Atoi(record[0])
		}
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			// ignore error: wrong number of fields in line
			if len(record) == 5 {
				cost, _ := strconv.ParseFloat(record[3], 64)
				status, _ := strconv.Atoi(record[4])
				response.Records = append(response.Records, DeliveryStatus{
					Name:     record[0],
					Mobile:   record[1],
					SendTime: record[2],
					Cost:     cost,
					Status:   StatusCode(status),
				})
			}
		}

		*v.(*DeliveryStatusResponse) = response

		return nil
	}

	result := new(DeliveryStatusResponse)
	_, err = c.Do(ctx, req, fn, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
