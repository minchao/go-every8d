package every8d

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestClient_GetCredit(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/API21/HTTP/getCredit.ashx", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, "88")
	})

	got, err := client.GetCredit(context.Background())
	if err != nil {
		t.Errorf("GetCredit returned unexpected error: %v", err)
	}
	if want := 88.0; got != want {
		t.Errorf("GetCredit returned %v, want %v", got, want)
	}
}

func TestClient_GetCredit_empty(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/API21/HTTP/getCredit.ashx", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, "")
	})

	_, err := client.GetCredit(context.Background())
	if err == nil {
		t.Error("Expected error response")
	}
}

func TestClient_GetCredit_invalid(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/API21/HTTP/getCredit.ashx", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, "invalid")
	})

	_, err := client.GetCredit(context.Background())
	if err == nil {
		t.Error("Expected error response")
	}
}
