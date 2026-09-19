package addtwonumbers

import (
	"reflect"
	"testing"

	"github.com/reftch/leetcode/internal/ds"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "example 1",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "example 2",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			name: "example 3",
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name: "different lengths with final carry",
			l1:   []int{9, 9},
			l2:   []int{1},
			want: []int{0, 0, 1},
		},
		{
			name: "single digit with carry",
			l1:   []int{5},
			l2:   []int{5},
			want: []int{0, 1},
		},
		{
			name: "carry chain across lists",
			l1:   []int{1, 9, 9},
			l2:   []int{9},
			want: []int{0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddTwoNumbers(ds.NewListNode(tt.l1), ds.NewListNode(tt.l2))
			if !reflect.DeepEqual(got.ToSlice(), tt.want) {
				t.Errorf("AddTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, got.ToSlice(), tt.want)
			}
		})
	}
}
