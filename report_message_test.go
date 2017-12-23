package every8d

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/google/go-querystring/query"
)

func TestParseReportMessage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	want := &ReportMessage{
		BatchID:      "00000000-0000-0000-0000-000000000000",
		Destination:  "+886987654321",
		ReportTime:   "20090210120000",
		StatusCode:   StatusCode(100),
		ReplyMessage: "Reply, Hello",
		MessageNo:    "001",
	}

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		got, err := ParseReportMessage(r)
		if err != nil {
			t.Errorf("ParseReportMessage returned unexpected error %v", err)
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Report message got %v, want %v", got, want)
		}
	})

	// Simulate the EVERY8D server report.
	q, _ := query.Values(want)
	u, _ := url.Parse("callback")
	u.RawQuery = q.Encode()

	req, _ := client.NewRequest(http.MethodGet, u.String(), nil)
	client.Do(context.Background(), req, nil, nil)
}
