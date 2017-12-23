package every8d

import (
	"testing"
)

func TestStatusCode_Text(t *testing.T) {
	tests := []struct {
		in  StatusCode
		out string
	}{
		{
			StatusCode(100),
			"發送成功",
		},
		{
			StatusCode(-100),
			"無此帳號。",
		},
		{
			StatusCode(100000),
			"Unknown StatusCode",
		},
	}

	for i, tt := range tests {
		if got, want := tt.in.Text(), tt.out; got != want {
			t.Errorf("Text %d. returned %v, want %v", i, got, want)
		}
	}
}
