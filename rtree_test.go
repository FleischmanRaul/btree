package btree

import "testing"

func hasEdge(rtree *RTree, key string) bool {
	_, ok := rtree.Edges[key]
	return ok
}

func TestCommonPrefix(t *testing.T) {
	tests := []struct {
		a, b, expected string
	}{
		{"apple", "application", "appl"},
		{"abc", "abc", "abc"},
		{"abc", "xyz", ""},
		{"", "abc", ""},
		{"abc", "", ""},
		{"a", "a", "a"},
	}

	for _, tt := range tests {
		result := commonPrefix(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("commonPrefix(%q, %q) = %q, want %q", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestBuildRTree_SingleValue(t *testing.T) {
	tree := buildRTree([]string{"apple"})
	if !hasEdge(tree, "apple") {
		t.Error("expected edge 'apple' at root")
	}
}

func TestBuildRTree_NoSharedPrefix(t *testing.T) {
	tree := buildRTree([]string{"apple", "banana"})
	if !hasEdge(tree, "apple") {
		t.Error("expected edge 'apple' at root")
	}
	if !hasEdge(tree, "banana") {
		t.Error("expected edge 'banana' at root")
	}
}

func TestBuildRTree_SharedPrefix(t *testing.T) {
	// "apple" and "application" share "appl"
	tree := buildRTree([]string{"apple", "application"})
	if !hasEdge(tree, "appl") {
		t.Error("expected shared prefix edge 'appl' at root")
	}
	appl := tree.Edges["appl"]
	if !hasEdge(appl, "e") {
		t.Error("expected edge 'e' under 'appl'")
	}
	if !hasEdge(appl, "ication") {
		t.Error("expected edge 'ication' under 'appl'")
	}
}

func TestBuildRTree_ExactPrefixMatch(t *testing.T) {
	// "ab" is an exact prefix of "abc"
	tree := buildRTree([]string{"ab", "abc"})
	if !hasEdge(tree, "ab") {
		t.Error("expected edge 'ab' at root")
	}
	ab := tree.Edges["ab"]
	if !hasEdge(ab, "c") {
		t.Error("expected edge 'c' under 'ab'")
	}
}

func TestBuildRTree_Duplicate(t *testing.T) {
	// duplicates should be silently skipped
	tree := buildRTree([]string{"apple", "apple"})
	if !hasEdge(tree, "apple") {
		t.Error("expected edge 'apple' at root")
	}
	if len(tree.Edges) != 1 {
		t.Errorf("expected 1 edge at root, got %d", len(tree.Edges))
	}
}

func TestBuildRTree_EmptyInput(t *testing.T) {
	tree := buildRTree([]string{})
	if len(tree.Edges) != 0 {
		t.Errorf("expected empty tree, got %d edges", len(tree.Edges))
	}
}

func TestBuildRTree_DeepNesting(t *testing.T) {
	// each value is a prefix of the next
	tree := buildRTree([]string{"a", "ab", "abc", "abcd"})
	if !hasEdge(tree, "a") {
		t.Error("expected edge 'a' at root")
	}
	a := tree.Edges["a"]
	if !hasEdge(a, "b") {
		t.Error("expected edge 'b' under 'a'")
	}
	ab := a.Edges["b"]
	if !hasEdge(ab, "c") {
		t.Error("expected edge 'c' under 'ab'")
	}
	abc := ab.Edges["c"]
	if !hasEdge(abc, "d") {
		t.Error("expected edge 'd' under 'abc'")
	}
}

func TestBuildRTree_MultipleSplits(t *testing.T) {
	// "ab", "abc", "abd" should split into "ab" -> {"c", "d"}
	tree := buildRTree([]string{"ab", "abc", "abd"})
	if !hasEdge(tree, "ab") {
		t.Error("expected edge 'ab' at root")
	}
	ab := tree.Edges["ab"]
	if !hasEdge(ab, "c") {
		t.Error("expected edge 'c' under 'ab'")
	}
	if !hasEdge(ab, "d") {
		t.Error("expected edge 'd' under 'ab'")
	}
}

func TestPPrint(t *testing.T) {
	tree := buildRTree([]string{"a", "ab", "abc", "abcd"})
	printRTree(tree)
	tree = buildRTree([]string{"ab", "abc", "abd", "apple", "application", "application2", "application3", "ag", "ag1", "ag23209743987234897432908732409873429087", "ag23209743987234897432908732409873429088", "ag3", "ag4"})
	printRTree(tree)
}
