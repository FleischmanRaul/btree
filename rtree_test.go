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

func TestContains_ExistingValues(t *testing.T) {
	tree := buildRTree([]string{"apple", "application", "banana"})
	for _, val := range []string{"apple", "application", "banana"} {
		if !tree.Contains(val) {
			t.Errorf("expected tree to contain %q", val)
		}
	}
}

func TestContains_NonExistingValues(t *testing.T) {
	tree := buildRTree([]string{"apple", "application", "banana"})
	for _, val := range []string{"app", "appl", "applications", "ban", "banan", "cherry"} {
		if tree.Contains(val) {
			t.Errorf("expected tree NOT to contain %q", val)
		}
	}
}

func TestContains_EmptyTree(t *testing.T) {
	tree := buildRTree([]string{})
	if tree.Contains("anything") {
		t.Error("expected empty tree to not contain anything")
	}
}

func TestContains_Prefix(t *testing.T) {
	// "ab" is inserted but "a" is not
	tree := buildRTree([]string{"ab", "abc"})
	if !tree.Contains("ab") {
		t.Error("expected tree to contain 'ab'")
	}
	if !tree.Contains("abc") {
		t.Error("expected tree to contain 'abc'")
	}
	if tree.Contains("a") {
		t.Error("expected tree NOT to contain 'a'")
	}
}

func TestContains_EmptyString(t *testing.T) {
	tree := buildRTree([]string{"apple"})
	if tree.Contains("") {
		t.Error("expected tree NOT to contain empty string")
	}
}

func TestContains_AfterInsert(t *testing.T) {
	tree := buildRTree([]string{})
	if tree.Contains("hello") {
		t.Error("expected tree NOT to contain 'hello' before insert")
	}
	tree.Insert("hello")
	if !tree.Contains("hello") {
		t.Error("expected tree to contain 'hello' after insert")
	}
}

func TestContains_Tricky(t *testing.T) {
	tree := buildRTree([]string{"apple1", "apple2", "apple"})
	if !tree.Contains("apple1") {
		t.Error("expected tree to contain 'hello' after insert")
	}
	if !tree.Contains("apple2") {
		t.Error("expected tree to contain 'apple2' after insert")
	}
	if !tree.Contains("apple") {
		t.Error("expected tree to contain 'apple' after insert")
	}
}

func TestContains_SharedPrefixPartial(t *testing.T) {
	// only full inserted words should match, not intermediate split nodes
	tree := buildRTree([]string{"apple", "application"})
	if tree.Contains("appl") {
		t.Error("expected tree NOT to contain 'appl' (shared prefix node)")
	}
	if tree.Contains("appli") {
		t.Error("expected tree NOT to contain 'appli'")
	}
}

func TestPPrint(t *testing.T) {
	tree := buildRTree([]string{"a", "ab", "abc", "abcd"})
	// printRTree(tree)
	tree = buildRTree([]string{"ab", "abc", "abd", "apple", "application", "application2", "application3", "ag", "ag1", "ag23209743987234897432908732409873429087", "ag23209743987234897432908732409873429088", "ag3", "ag4"})
	// printRTree(tree)
	tree = buildRTree([]string{"bea", "berti", "beke", "andor", "akvarium"})
	// printRTree(tree)
	tree = buildRTree([]string{})
	tree.Insert("first")
	tree.Insert("fisst")
	tree.Insert("fsss")
	// printRTree(tree)
}
