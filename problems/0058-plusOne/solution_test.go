package plusOne

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 4},
		},
		{
			name: "example 2",
			nums: []int{4, 3, 2, 1},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "example 3",
			nums: []int{9},
			want: []int{1, 0},
		},
		{
			name: "example 4",
			nums: []int{9, 8, 9},
			want: []int{9, 9, 0},
		},
		{
			name: "example 5",
			nums: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
			want: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("onePlus() = %v, want %v", got, tt.want)
			}
		})
	}
}
