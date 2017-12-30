package every8d

import (
	"context"
	"encoding/csv"
	"io"
	"strconv"

	"github.com/go-playground/form"
)

// Message represents an SMS object.
type Message struct {
	// Message title. Empty titles are accepted.
	// The title will not be sent with the SMS; it is just a note.
	Subject string `form:"SB,omitempty"`

	// Message content.
	Content string `form:"MSG"`

	// Receiver's mobile number.
	// Format: +88612345678 or 0912345678
	// Separator: (,) e.g. 0912345678,0922333444
	Destination string `form:"DEST"`

	// Reservation time
	// Send immediately: No input (empty).
	// Reservation send:Please input the reservation time, using this format: yyyyMMddHHmnss, e.g. 20090131153000
	ReservationTime string `form:"ST,omitempty"`

	// SMS validity period of unit: minutes.
	// if not specified, then the platform default validity period is 1440 minutes.
	RetryTime int `form:"RETRYTIME,omitempty"`

	// Message record no.
	MessageNo string `form:"MR,omitempty"`
}

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

// SendMMS sends a MMS.
func (c *Client) SendMMS(ctx context.Context, message MMS) (*SendResponse, error) {
	return c.send(ctx, "API21/HTTP/MMS/sendMMS.ashx", message)
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
