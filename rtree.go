package btree

import (
	"fmt"
	"sort"
)

// Definition of a radix tree.
type RTree struct {
	Metadata string
	Final    bool
	Edges    map[string]*RTree
}

func newNode() *RTree {
	return &RTree{
		Edges: make(map[string]*RTree),
	}
}

func BuildRTree(vals []string) *RTree {
	head := newNode()

	for _, val := range vals {
		head.Insert(val)
	}
	return head
}

func navigate(val string, rtree *RTree) (string, *RTree) {
	for k := range rtree.Edges {
		cp := commonPrefix(val, k)
		if len(cp) == len(k) {
			return navigate(val[len(cp):], rtree.Edges[k])
		}

		if len(cp) > 0 {
			rtree.Edges[cp] = newNode()
			rtree.Edges[cp].Edges[k[len(cp):]] = rtree.Edges[k]
			delete(rtree.Edges, k)
			return val[len(cp):], rtree.Edges[cp]
		}
	}
	return val, rtree
}

func navigate2(val string, rtree *RTree) (string, *RTree, string, *RTree, string, *RTree) {
	return navigate2helper(val, rtree, "", nil, "", nil)
}

func navigate2helper(val string, rtree *RTree, valp string, rtreep *RTree, valg string, rtreeg *RTree) (string, *RTree, string, *RTree, string, *RTree) {
	for k := range rtree.Edges {
		cp := commonPrefix(val, k)
		if len(cp) == len(k) {
			return navigate2helper(val[len(cp):], rtree.Edges[k], k, rtree, valp, rtreep)
		}

		if len(cp) > 0 {
			rtree.Edges[cp] = newNode()
			rtree.Edges[cp].Edges[k[len(cp):]] = rtree.Edges[k]
			delete(rtree.Edges, k)
			return val[len(cp):], rtree.Edges[cp], cp, rtree, valp, rtreep
		}
	}
	return val, rtree, valp, rtreep, valg, rtreeg
}

func commonPrefix(a, b string) string {
	i := 0
	for ; i < len(a) && i < len(b) && a[i] == b[i]; i++ {
	}
	return a[:i]
}

func (rtree *RTree) PrintTree() {
	printHelper(rtree, "", "", "root", false)
}

func printHelper(node *RTree, prefix string, connector string, val string, ins bool) {
	fmt.Println(prefix + connector + val)
	keys_list := []string{}
	for k := range node.Edges {
		keys_list = append(keys_list, k)
	}
	sort.Strings(keys_list)

	newPrefix := prefix + "    "
	if ins {
		newPrefix = prefix + "│   "
	}

	for i, k := range keys_list {
		if i == len(keys_list)-1 {
			printHelper(node.Edges[k], newPrefix, "└── ", k, false)
		} else {
			printHelper(node.Edges[k], newPrefix, "├── ", k, true)
		}
	}
}

func (rtree *RTree) Insert(val string) {
	v, r := navigate(val, rtree)
	if len(v) == 0 {
		r.Final = true
		return
	}
	r.Edges[v] = &RTree{
		Metadata: "",
		Final:    true,
		Edges:    make(map[string]*RTree),
	}
}

func (rtree *RTree) Delete(val string) bool {
	if rtree == nil {
		return false
	}

	if !rtree.Contains(val) {
		return false
	}

	_, r, pv, p, gv, g := navigate2(val, rtree)
	r.Final = false

	for k := range r.Edges {
		p.Edges[pv+k] = r.Edges[k]
	}
	delete(p.Edges, pv)

	if len(p.Edges) == 1 {
		for k := range p.Edges {
			g.Edges[gv+k] = p.Edges[k]
		}
		delete(g.Edges, gv)
	}

	return true
}

func (rtree *RTree) Contains(val string) bool {
	v, r := navigate(val, rtree)
	return len(v) == 0 && r.Final
}

func (rtree *RTree) PartialContains(val string) bool {
	v, _ := navigate(val, rtree)
	return len(v) == 0
}

func (rtree *RTree) List() []string {
	return listHelper(rtree, "")
}

func listHelper(node *RTree, acc string) []string {
	result := []string{}
	if node.Final {
		result = append(result, acc)
	}
	for k, v := range node.Edges {
		result = append(result, listHelper(v, acc+k)...)
	}
	return result
}

// func ForEachPrefix() {
//  TODO
// }

// func DeletePrefix() {}
// func Iterator() Iterator

// func New() {}
// func Len() {}
