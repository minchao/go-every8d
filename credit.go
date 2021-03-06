package every8d

import (
	"context"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
)

// GetCredit retrieves your account balance.
func (c *Client) GetCredit(ctx context.Context) (float64, error) {
	req, err := c.NewFormRequest("API21/HTTP/getCredit.ashx", url.Values{})
	if err != nil {
		return 0.0, err
	}

	fn := func(body io.Reader, v interface{}) error {
		result, _ := ioutil.ReadAll(body)
		credit, err := strconv.ParseFloat(string(result), 64)
		if err != nil {
			return err
		}

		*v.(*float64) = credit

		return nil
	}

	credit := new(float64)
	_, err = c.Do(ctx, req, fn, credit)
	if err != nil {
		return 0.0, err
	}

	return *credit, nil
}
