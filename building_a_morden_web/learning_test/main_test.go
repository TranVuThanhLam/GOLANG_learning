package main

import "testing"

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0, true},
	{"test-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivide(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divident, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expect an err")
			}
		} else {
			if err != nil {
				t.Error("not expect an err")
			}
		}
		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}

// maybe la phai dung goland moi dung dc lenh nay
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out
