package simplifyPath

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "example 1 trailing slash removed",
			path: "/home/",
			want: "/home",
		},
		{
			name: "example 2 multiple slashes collapsed",
			path: "/home//foo/",
			want: "/home/foo",
		},
		{
			name: "example 3 double dot goes to parent",
			path: "/home/user/Documents/../Pictures",
			want: "/home/user/Pictures",
		},
		{
			name: "example 4 cannot go above root",
			path: "/../",
			want: "/",
		},
		{
			name: "example 5 triple dot is valid name",
			path: "/.../a/../b/c/../d/./",
			want: "/.../b/d",
		},
		{
			name: "example 6",
			path: "/a/./b/../../c/",
			want: "/c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("simplifyPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}
