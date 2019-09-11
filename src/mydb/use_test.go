package mydb

import "testing"

func TestGetOne1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"sf",},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetOne1()
		})
	}
}
