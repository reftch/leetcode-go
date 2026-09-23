package climbStairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "one step",
			n:    1,
			want: 1,
		},
		{
			name: "example 1",
			n:    2,
			want: 2,
		},
		{
			name: "example 2",
			n:    3,
			want: 3,
		},
		{
			name: "four steps",
			n:    4,
			want: 5,
		},
		{
			name: "five steps",
			n:    5,
			want: 8,
		},
		{
			name: "ten steps",
			n:    10,
			want: 89,
		},
		{
			name: "twenty steps",
			n:    20,
			want: 10946,
		},
		{
			name: "max constraint",
			n:    45,
			want: 1836311903,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
