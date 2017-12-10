package every8d

import (
	"context"
	"encoding/csv"
	"io"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
)

// SendResponse represents the response of send an SMS.
type SendResponse struct {
	// Balance credit.
	// Negative means there was a delivery failure and the system can't process this command.
	Credit float64

	// Sent messages.
	Sent int

	// This shows spent points.
	Cost float64

	// Unsent messages with no credit charged.
	Unsent int

	// Batch ID. e.g. 220478cc-8506-49b2-93b7-2505f651c12e
	BatchID string
}

// Send sends an SMS.
func (c *Client) Send(ctx context.Context, message Message) (*SendResponse, error) {
	q, _ := query.Values(message)
	u, _ := url.Parse("API21/HTTP/sendSMS.ashx")
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	fn := func(body io.Reader, v interface{}) error {
		r := csv.NewReader(body)
		record, err := r.Read()
		if err != nil {
			return err
		}

		credit, _ := strconv.ParseFloat(record[0], 64)
		sent, _ := strconv.Atoi(record[1])
		cost, _ := strconv.ParseFloat(record[2], 64)
		unsent, _ := strconv.Atoi(record[3])

		*v.(*SendResponse) = SendResponse{
			credit,
			sent,
			cost,
			unsent,
			record[4],
		}

		return nil
	}

	result := &SendResponse{}
	_, err = c.Do(ctx, req, fn, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
