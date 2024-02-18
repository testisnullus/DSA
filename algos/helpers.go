package algos

type TreeNode struct {
	Value    int
	Children []*TreeNode
}

type Stack struct {
	storage []*TreeNode
}

type Queue struct {
	storage []*TreeNode
}

func NewStack(s *TreeNode) *Stack {
	return &Stack{
		storage: []*TreeNode{s},
	}
}

func NewQueue(s *TreeNode) *Queue {
	return &Queue{
		storage: []*TreeNode{s},
	}
}

func (s *Stack) Push(value *TreeNode) {
	s.storage = append(s.storage, value)
}

func (s *Stack) Pop() *TreeNode {
	current := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]

	return current
}

func (q *Queue) Enqueue(value *TreeNode) {
	q.storage = append(q.storage, value)
}

func (q *Queue) Dequeue() *TreeNode {
	current := q.storage[0]
	q.storage = q.storage[1:]

	return current
}
