package longestCommonPrefix

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "example 1",
			input: []string{"flower", "flow", "flight"},
			want:  "fl",
		},
		{
			name:  "example 2",
			input: []string{"dog", "racecar", "car"},
			want:  "",
		},
		{
			name:  "empty prefix",
			input: []string{"a", "b", "c"},
			want:  "",
		},
		{
			name:  "single string",
			input: []string{"abc"},
			want:  "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
