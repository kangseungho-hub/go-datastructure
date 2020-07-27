package datastructure

import (
	"fmt"
	"reflect"
)

type Stack struct {
	tail *StackNode
}

type StackNode struct {
	data interface{}
	prev *StackNode
}

func NewStack() *Stack {
	return &Stack{tail: nil}
}

func NewStackNode(data interface{}) *StackNode {
	return &StackNode{data: data, prev: nil}
}

func (s *Stack) AppendStack(data interface{}) bool {

	if reflect.ValueOf(data).IsNil() {
		return false
	}

	newNode := NewStackNode(data)
	if s.tail == nil {
		s.tail = newNode
	} else {
		newNode.prev = s.tail
		s.tail = newNode
	}

	return true
}

func (s *Stack) Pop() interface{} {
	data := s.tail.data
	s.tail = s.tail.prev
	return data
}

func (s *Stack) StackIsEmpty() bool {
	if s.tail == nil {
		return true
	}
	return false
}

func (s *Stack) PrintStack() {
	currentNode := s.tail

	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.prev
	}
}
