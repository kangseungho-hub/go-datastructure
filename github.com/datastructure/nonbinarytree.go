package datastructure

// import "fmt"

// type TreeNode struct {
// 	Val    int
// 	Childs []*TreeNode
// }

// type Tree struct {
// 	Root *TreeNode
// }

// func (t *TreeNode) AddNode(val int) {
// 	t.Childs = append(t.Childs, &TreeNode{Val: val})
// }
// func (t *Tree) AddNode(val int) {
// 	if t.Root == nil {
// 		t.Root = &TreeNode{Val: val}
// 	} else {
// 		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
// 	}

// }
// func (t *Tree) RecursiveDFS() {
// 	RecursiveDFS(t.Root)
// }
// func RecursiveDFS(n *TreeNode) {
// 	fmt.Println(n.Val)

// 	for _, child := range n.Childs {
// 		RecursiveDFS(child)
// 	}
// }
