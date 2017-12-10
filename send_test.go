package every8d

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_Send(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/API21/HTTP/sendSMS.ashx", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, "87.00,1,1,0,00000000-0000-0000-0000-000000000000")
	})

	message := Message{
		Subject:         "note",
		Content:         "Hello, 世界",
		Destination:     "+88612345678",
		ReservationTime: "",
		RetryTime:       0,
	}

	want := &SendResponse{
		Credit:  87.0,
		Sent:    1,
		Cost:    1,
		Unsent:  0,
		BatchID: "00000000-0000-0000-0000-000000000000",
	}

	got, err := client.Send(context.Background(), message)
	if err != nil {
		t.Errorf("Send returned unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Send returned %+v, want %+v", got, want)
	}
}
