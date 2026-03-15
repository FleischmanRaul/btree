package btree

import (
	"fmt"
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
	fmt.Println(rtree.Edges)
	for key := range rtree.Edges {
		printRTree(rtree.Edges[key])
	}
}

func prettyPrintRTree(rtree *RTree) {
	
}