package algos

import "fmt"

type TreeNode struct {
	Value    int
	Children []*TreeNode
}

func DFSRecursive(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	for _, child := range node.Children {
		DFSRecursive(child)
	}
}

func DFSStack(node *TreeNode) {
	if node == nil {
		return
	}

	stack := []*TreeNode{node}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println(current.Value)
		for i := len(current.Children) - 1; i >= 0; i-- {
			stack = append(stack, current.Children[i])
		}
	}
}
