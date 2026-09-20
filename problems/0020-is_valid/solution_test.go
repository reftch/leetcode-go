package is_valid

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "example 1",
			input: "()",
			want:  true,
		},
		{
			name:  "example 2",
			input: "()[]{}",
			want:  true,
		},
		{
			name:  "example 3",
			input: "(]",
			want:  false,
		},
		{
			name:  "example 4",
			input: "([])",
			want:  true,
		},
		{
			name:  "example 5",
			input: "([)]",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is_valid(tt.input); got != tt.want {
				t.Errorf("is_valid(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
