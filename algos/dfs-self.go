package algos

import "fmt"

func DFSSelf(node *TreeNode) {
	if node == nil {
		return
	}

	stack := []*TreeNode{node}

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println(current.Value)
		for i := len(current.Children) - 1; i >= 0; i-- {
			stack = append(stack, current.Children[i])
		}
	}
}
