package datastructure

import (
	"fmt"
	"reflect"
)

type queue struct {
	head *QNode
	tail *QNode
}

type QNode struct {
	data interface{}
	next *QNode
}

func NewQueue() *queue {
	return &queue{head: nil, tail: nil}
}

func NewQNode(data interface{}) *QNode {
	return &QNode{data: data, next: nil}
}

func (q *queue) Add(data interface{}) bool {

	if reflect.ValueOf(data).IsNil() {
		return false
	}
	newQNode := NewQNode(data)
	if q.IsEmpty() {
		q.tail = newQNode
		q.head = q.tail
	} else {
		q.tail.next = newQNode
		q.tail = newQNode
	}

	return true
}

func (q *queue) IsEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}

func (q *queue) PrintQueue() {
	currentQNode := q.head
	for currentQNode != nil {
		fmt.Println(currentQNode.data)
		currentQNode = currentQNode.next
	}
}

func (q *queue) Poll() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := q.Peek()
	q.ReMove()
	return data
}

func (q *queue) ReMove() {
	q.head = q.head.next
}

func (q *queue) Peek() interface{} {
	return q.head.data
}
