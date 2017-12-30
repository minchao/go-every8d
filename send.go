package every8d

import (
	"context"
	"encoding/csv"
	"io"
	"strconv"

	"github.com/go-playground/form"
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
	return c.send(ctx, "API21/HTTP/sendSMS.ashx", message)
}

// MMS represents an MMS object.
type MMS struct {
	Message

	// Image file, binary base64 encoded.
	Attachment string `form:"ATTACHMENT"`

	// Image file extension, support jpg/jpeg/png/git.
	Type string `form:"TYPE"`
}

func (c *Client) send(ctx context.Context, urlStr string, message interface{}) (*SendResponse, error) {
	f, _ := form.NewEncoder().Encode(message)

	req, err := c.NewFormRequest(urlStr, f)
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
