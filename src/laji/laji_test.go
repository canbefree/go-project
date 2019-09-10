package laji

import "testing"

func Test_what_point_do(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name : "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			what_point_do()
		})
	}
}
