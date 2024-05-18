package btree

import (
	"fmt"
	"testing"
)

func TestBuildBTree(t *testing.T) {
	nums := []int{3}
	result := buildBTree(nums)
	fmt.Println(nums)
	printBTree(result)

	nums = []int{3, 4, 5}
	result = buildBTree(nums)
	fmt.Println(nums)
	printBTree(result)

	nums = []int{3, 4, 5, 6, 7}
	result = buildBTree(nums)
	fmt.Println(nums)
	printBTree(result)

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result = buildBTree(nums)
	fmt.Println(nums)
	printBTree(result)
}

func TestMinBTree(t *testing.T) {
	nums := []int{3}
	result := buildBTree(nums)
	printBTree(result)
	fmt.Println(minBTree(result))

	nums = []int{3, 1}
	result = buildBTree(nums)
	printBTree(result)
	fmt.Println(minBTree(result))

	nums = []int{3, 1, 2}
	result = buildBTree(nums)
	printBTree(result)
	fmt.Println(minBTree(result))

	nums = []int{3, 1, -2}
	result = buildBTree(nums)
	printBTree(result)
	fmt.Println(minBTree(result))

	nums = []int{3, 1, -2, 4, 5, 6, 7, -9, 10, 2, 3, 1, 1, 0, 0, 0}
	result = buildBTree(nums)
	printBTree(result)
	fmt.Println(minBTree(result))

	result = nil
	printBTree(result)
	fmt.Println(minBTree(result))
}
