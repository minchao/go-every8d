package every8d

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/go-playground/form"
)

func TestMessage_toURLValues(t *testing.T) {
	tests := []struct {
		in   Message
		want url.Values
	}{
		{
			Message{
				Content:     "Hello",
				Destination: "+886987654321",
			},
			url.Values{
				"MSG":  {"Hello"},
				"DEST": {"+886987654321"},
			},
		},
		{
			Message{},
			url.Values{
				"MSG":  {""},
				"DEST": {""},
			},
		},
		{
			Message{
				Subject:         "Subject",
				Content:         "Hello",
				Destination:     "+886987654321",
				ReservationTime: "20090131153000",
				RetryTime:       3600,
				MessageNo:       "001",
			},
			url.Values{
				"SB":        {"Subject"},
				"MSG":       {"Hello"},
				"DEST":      {"+886987654321"},
				"ST":        {"20090131153000"},
				"RETRYTIME": {"3600"},
				"MR":        {"001"},
			},
		},
	}

	for i, tt := range tests {
		q, _ := form.NewEncoder().Encode(tt.in)

		if got, want := q, tt.want; !reflect.DeepEqual(got, want) {
			t.Errorf("Message to url.Values %d. returned %v, want %v", i, got, want)
		}
	}
}
