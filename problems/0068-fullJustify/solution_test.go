package fullJustify

import "testing"

func TestFullJustify(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		width int
		want  []string
	}{
		{
			name:  "example 1",
			words: []string{"This", "is", "an", "example", "of", "text", "justification."},
			width: 16,
			want:  []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			name:  "example 2",
			words: []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			width: 16,
			want:  []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			name:  "example 3",
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			width: 20,
			want:  []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
		{
			name:  "example 5",
			words: []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"},
			width: 16,
			want:  []string{"ask   not   what", "your country can", "do  for  you ask", "what  you can do", "for your country"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.words, tt.width); !compareArrays(got, tt.want) {
				t.Errorf("fullJustify(%s, %d) = %v, want %v", tt.words, tt.width, got, tt.want)
			}
		})
	}
}

func compareArrays(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
