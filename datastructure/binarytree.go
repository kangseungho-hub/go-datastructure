package datastructure

import (
	"fmt"
)

type Tree struct {
	Root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
	count int
}

//NewTree is creator function of Tree structure
func NewTree() *Tree {
	return &Tree{Root: nil}
}

func NewNode(data int) *Node {
	return &Node{data: data, left: nil, right: nil, count: 0}
}

//Insert is append element function to Tree
func (t *Tree) Insert(data int) {
	// crueentNode := *t.Root 이거는 그 값을 복사하는 거자너~
	//지금 여기 값이 안바뀌는게 문제임, 이거 튜토리얼 영상 보면될듯?
	//이유 : currentNode를 바꾸면 그냥 새 노드 포인터로 교체하는 거니까 당연히 안되는 거고
	//root에 추가할때를 따로해서 root를 current노드로 한다음 .left .right이렇게 접근해야
	//새로운 노드 포인터를 그 property에다가 넣을 수 있는 거
	if t.IsEmpty() {
		t.Root = NewNode(data)
	} else {
		currentNode := t.Root
		for true {
			if currentNode.data > data {
				if currentNode.left == nil {
					currentNode.left = NewNode(data)
					break
				} else {
					currentNode = currentNode.left
				}
			} else if currentNode.data < data {
				if currentNode.right == nil {
					currentNode.right = NewNode(data)
					break
				} else {
					currentNode = currentNode.right
				}
			} else if currentNode.data == data {
				currentNode.count++
				break
			}
		}
	}

}

//IsEmpty function will check tree is empty
func (t *Tree) IsEmpty() bool {
	if t.Root == nil {
		return true
	}
	return false
}

func (t *Tree) RecurSiveDFS() {
	RecurSiveDFS(t.Root)
}

func RecurSiveDFS(n *Node) {
	if n != nil {
		fmt.Printf("%d(%d) -> ", n.data, n.count)
		RecurSiveDFS(n.left)
		RecurSiveDFS(n.right)
	}
}

func (t *Tree) StackDFS() {
	fmt.Println("")
	stack := NewStack()
	currentNode := t.Root

	stack.AppendStack(currentNode)
	for !stack.StackIsEmpty() {

		popedNode := stack.Pop()
		fmt.Printf("%d(%d) -> ", popedNode.(*Node).data, popedNode.(*Node).count)

		stack.AppendStack(popedNode.(*Node).right)
		stack.AppendStack(popedNode.(*Node).left)
	}

}

func (t *Tree) BFS() {
	Q := NewQueue()

	currentNode := t.Root
	Q.Add(currentNode)

	fmt.Println("")

	for !Q.IsEmpty() {
		polledNode := Q.Poll()
		fmt.Printf("%d ->", polledNode.(*Node).data)
		Q.Add(polledNode.(*Node).left)
		Q.Add(polledNode.(*Node).right)
	}

}

func PrintChild(n *Node) {
	fmt.Println(n.left)
	fmt.Println(n.right)
}
