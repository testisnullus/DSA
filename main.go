package main

import (
	"fmt"
	"github.com/testisnullus/data-structs-and-algorithms/algos"
)

var sortedArr = []int{1, 5, 12, 23, 46, 87, 101, 187, 291, 332, 341, 396, 401, 502, 505, 555, 591, 600}

// Example tree:
//
//	     1
//	   / | \
//	  2  3  4
//	 / \    |
//	5   6   7
//	   /
//	  8
var tree = &algos.TreeNode{
	Value: 1,
	Children: []*algos.TreeNode{
		{2,
			[]*algos.TreeNode{
				{5,
					nil,
				},
				{6,
					[]*algos.TreeNode{
						{8,
							nil},
					}},
			}},
		{3,
			nil},
		{4,
			[]*algos.TreeNode{
				{7,
					nil},
			}},
	}}

func main() {
	fmt.Printf("BS. Index of elem: %d\n", algos.BinarySearch(sortedArr, 101))

	fmt.Println("DFS Recursive. Traversed tree:")
	algos.DFSRecursive(tree)

	fmt.Println("DFS Stack. Traversed tree:")
	algos.DFSStack(tree)

	fmt.Println("BFS. Traversed tree:")
	algos.BFS(tree)
}
