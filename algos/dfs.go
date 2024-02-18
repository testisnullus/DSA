package algos

import "fmt"

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

	stack := NewStack(node)

	for len(stack.storage) > 0 {
		current := stack.Pop()

		fmt.Println(current.Value)
		for i := len(current.Children) - 1; i >= 0; i-- {
			stack.Push(current.Children[i])
		}
	}
}
