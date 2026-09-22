package strStr

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		target string
		want   int
	}{
		{
			name:   "example 1",
			input:  "sadbutsad",
			target: "sad",
			want:   0,
		},
		{
			name:   "example 2",
			input:  "leetcode",
			target: "leeto",
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.input, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
