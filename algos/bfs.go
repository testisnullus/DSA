package algos

import "fmt"

func BFS(node *TreeNode) {
	if node == nil {
		return
	}

	queue := NewQueue(node)
	for len(queue.storage) != 0 {
		current := queue.Dequeue()

		fmt.Println(current.Value)
		for _, child := range current.Children {
			queue.Enqueue(child)
		}
	}
}
