package lengthOfLastWord

import (
	"reflect"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example 1",
			input: "Hello World",
			want:  5,
		},
		{
			name:  "example 2",
			input: "   fly me   to   the moon  ",
			want:  4,
		},
		{
			name:  "example 3",
			input: "luffy is still joyboy",
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
