package every8d

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_Send(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/API21/HTTP/sendSMS.ashx", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{
			"SB":   "note",
			"MSG":  "Hello, 世界",
			"DEST": "+88612345678",
		})
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

func TestClient_Send_unknownError(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/API21/HTTP/sendSMS.ashx", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, "-99, 主機端發生不明錯誤，請與廠商窗口聯繫。")
	})

	_, err := client.Send(context.Background(), Message{})
	if err == nil {
		t.Fatal("Expected error to be returned.")
	}
	if got, want := StatusCode(-99), err.(*ErrorResponse).ErrorCode; got != want {
		t.Errorf("Error = %v, want %v", got, want)
	}
}
