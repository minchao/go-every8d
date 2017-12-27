package every8d

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_GetDeliveryStatus(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/API21/HTTP/getDeliveryStatus.ashx", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"BID": "00000000-0000-0000-0000-000000000000",
			"PNO": "1",
		})
		fmt.Fprint(w, `2
Test	+886987654321	2017/12/18 23:14:17	1	100
	+886987654321	2017/12/18 23:14:18	0	101`)
	})

	want := &DeliveryStatusResponse{
		Count: 2,
		Records: []DeliveryStatus{
			{
				Name:     "Test",
				Mobile:   "+886987654321",
				SendTime: "2017/12/18 23:14:17",
				Cost:     1,
				Status:   StatusCode(100),
			},
			{
				Name:     "",
				Mobile:   "+886987654321",
				SendTime: "2017/12/18 23:14:18",
				Cost:     0,
				Status:   StatusCode(101),
			},
		},
	}

	got, err := client.GetDeliveryStatus(context.Background(), "00000000-0000-0000-0000-000000000000", "1")
	if err != nil {
		t.Errorf("GetDeliveryStatus returned unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetDeliveryStatus returned %+v, want %+v", got, want)
	}
}

func TestClient_GetMMSDeliveryStatus(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/API21/HTTP/MMS/getDeliveryStatus.ashx", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"BID": "00000000-0000-0000-0000-000000000000",
			"PNO": "1",
		})
		fmt.Fprint(w, `2
Test	+886987654321	2017/12/18 23:14:17	1	100
	+886987654321	2017/12/18 23:14:18	0	101`)
	})

	want := &DeliveryStatusResponse{
		Count: 2,
		Records: []DeliveryStatus{
			{
				Name:     "Test",
				Mobile:   "+886987654321",
				SendTime: "2017/12/18 23:14:17",
				Cost:     1,
				Status:   StatusCode(100),
			},
			{
				Name:     "",
				Mobile:   "+886987654321",
				SendTime: "2017/12/18 23:14:18",
				Cost:     0,
				Status:   StatusCode(101),
			},
		},
	}

	got, err := client.GetMMSDeliveryStatus(context.Background(), "00000000-0000-0000-0000-000000000000", "1")
	if err != nil {
		t.Errorf("GetMMSDeliveryStatus returned unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetMMSDeliveryStatus returned %+v, want %+v", got, want)
	}
}
