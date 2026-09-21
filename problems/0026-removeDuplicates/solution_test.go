package removeDuplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "example 1", input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},
		{name: "example 1", input: []int{1, 1, 2}, want: 2},
		{name: "all same", input: []int{1, 1, 1, 1}, want: 1},
		{name: "no dups", input: []int{1, 2, 3, 4, 5}, want: 5},
		{name: "empty", input: []int{}, want: 0},
		{name: "single", input: []int{7}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.input); got != tt.want {
				t.Errorf("removeDuplicates(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
