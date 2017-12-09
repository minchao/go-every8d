package every8d

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestClient_GetCredit(t *testing.T) {
	setup()
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
