package addBinary

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		input2 string
		want   string
	}{
		{
			name:   "example 1",
			input1: "11",
			input2: "1",
			want:   "100",
		},
		{
			name:   "example 2",
			input1: "1010",
			input2: "1011",
			want:   "10101",
		},
		{
			name:   "example 3",
			input1: "0",
			input2: "1",
			want:   "1",
		},
		{
			name:   "example 4",
			input1: "1111",
			input2: "1111",
			want:   "11110",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.input1, tt.input2); got != tt.want {
				t.Errorf("addBinary(%s, %s) = %v, want %v", tt.input1, tt.input2, got, tt.want)
			}
		})
	}
}
