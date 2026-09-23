package mySqrt

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "example 1",
			x:    4,
			want: 2,
		},
		{
			name: "example 2",
			x:    8,
			want: 2,
		},
		{
			name: "zero",
			x:    0,
			want: 0,
		},
		{
			name: "one",
			x:    1,
			want: 1,
		},
		{
			name: "two rounds down to one",
			x:    2,
			want: 1,
		},
		{
			name: "three rounds down to one",
			x:    3,
			want: 1,
		},
		{
			name: "perfect square nine",
			x:    9,
			want: 3,
		},
		{
			name: "non-perfect square fifteen",
			x:    15,
			want: 3,
		},
		{
			name: "perfect square sixteen",
			x:    16,
			want: 4,
		},
		{
			name: "large non-perfect square",
			x:    2147395599,
			want: 46339,
		},
		{
			name: "max int32",
			x:    2147483647,
			want: 46340,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt(%v) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
