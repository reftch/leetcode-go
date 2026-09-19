package ds

import (
	"reflect"
	"testing"
)

func TestListNodeRoundTrip(t *testing.T) {
	head := NewListNode([]int{1, 2, 3})
	if got := head.ToSlice(); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("ToSlice() = %v, want [1 2 3]", got)
	}
	if NewListNode(nil) != nil {
		t.Fatal("NewListNode(nil) should return nil")
	}
}

func TestTreeNodeRoundTrip(t *testing.T) {
	root := NewTreeNode([]any{1, 2, 3, nil, 4})
	want := []any{1, 2, 3, nil, 4}
	if got := root.ToLevelOrder(); !reflect.DeepEqual(got, want) {
		t.Fatalf("ToLevelOrder() = %v, want %v", got, want)
	}
	if NewTreeNode(nil) != nil {
		t.Fatal("NewTreeNode(nil) should return nil")
	}
}
