package btree

import "fmt"

// Definition of a binary tree.
type BTree struct {
	Val   int
	Left  *BTree
	Right *BTree
}

func buildBTree(nums []int) *BTree {
	if len(nums) == 0 {
		return nil
	}

	head := &BTree{
		Val:   nums[0],
		Left:  nil,
		Right: nil,
	}
	current := []*BTree{head}

	for i := 1; i < len(nums); {
		next_row := []*BTree{}
		for j := 0; j < len(current); j++ {
			// add to left
			current[j].Left = &BTree{
				Val:   nums[i],
				Left:  nil,
				Right: nil,
			}

			i++
			if i >= len(nums) {
				return head
			}

			current[j].Right = &BTree{
				Val:   nums[i],
				Left:  nil,
				Right: nil,
			}
			i++
			if i >= len(nums) {
				return head
			}

			next_row = append(next_row, current[j].Left, current[j].Right)
		}
		current = next_row
	}
	return head
}

func printBTree(btree *BTree) {
	if btree == nil {
		return
	}
	fmt.Println(btree.Val)
	printBTree(btree.Left)
	printBTree(btree.Right)
}

func prettyPrintBTree(btree *BTree) {
	fmt.Print("Unimplemented")
}

func dfs(btree *BTree) {
	// Depth-first search
	fmt.Print("Unimplemented")
}

func bfs(btree *BTree) {
	// Breath-first search
	fmt.Print("Unimplemented")
}

func minBTree(btree *BTree) (int, bool) {
	// Find smallest element
	if btree == nil {
		return 0, true
	}

	if btree.Left == nil && btree.Right == nil {
		return btree.Val, false
	}

	if btree.Left != nil && btree.Right == nil {
		l, _ := minBTree(btree.Left)
		return min(btree.Val, l), false
	}

	if btree.Left == nil && btree.Right != nil {
		r, _ := minBTree(btree.Right)
		return min(btree.Val, r), false
	}

	l, _ := minBTree(btree.Left)
	r, _ := minBTree(btree.Right)
	return min(btree.Val, l, r), false
}

func maxBTree(btree *BTree) {
	// Find largest element
	fmt.Print("Unimplemented")
}

func insert(btree *BTree) {
	// Insert a new element
	fmt.Print("Unimplemented")
}

func remove(btree *BTree) {
	// Delete element if it exists
	fmt.Print("Unimplemented")
}

/*
1. Create a BTree interface
2. Use generic type instead of int
3. Implement multiple types of binary trees: Balanced, sorted, simple.
*/
