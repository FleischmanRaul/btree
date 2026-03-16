package btree

import (
	"fmt"
	"sort"
)

// Definition of a radix tree.
type RTree struct {
	Metadata string
	Edges    map[string]*RTree
}

func buildRTree(vals []string) *RTree {
	head := &RTree{
		Metadata: "",
		Edges:    make(map[string]*RTree),
	}

	for _, val := range vals {
		v, r := navigate(val, head)

		if len(v) == 0 {
			fmt.Println("Duplicate: ", val)
			continue
		}

		r.Edges[v] = &RTree{
			Metadata: "",
			Edges:    make(map[string]*RTree),
		}

	}
	return head
}

func navigate(val string, rtree *RTree) (string, *RTree) {
	for i := len(val); i > 0; i-- {
		cont, ok := rtree.Edges[val[0:i]]
		if ok {
			return navigate(val[i:], cont)
		}
	}

	// partial match
	for k := range rtree.Edges {
		cp := commonPrefix(val, k)
		if len(cp) > 0 {
			rtree.Edges[cp] = &RTree{
				Metadata: "",
				Edges:    make(map[string]*RTree),
			}
			rtree.Edges[cp].Edges[k[len(cp):]] = rtree.Edges[k]
			delete(rtree.Edges, k)
			return val[len(cp):], rtree.Edges[cp]
		}
	}

	return val, rtree
}

func commonPrefix(a, b string) string {
	i := 0
	for ; i < len(a) && i < len(b) && a[i] == b[i]; i++ {
	}
	return a[:i]
}

func printRTree(rtree *RTree) {
	printRTreeInt(rtree, "", "", "root", false)
}

func printRTreeInt(node *RTree, prefix string, connector string, val string, ins bool) {
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
			printRTreeInt(node.Edges[k], newPrefix, "└── ", k, false)
		} else {
			printRTreeInt(node.Edges[k], newPrefix, "├── ", k, true)
		}
	}
}
