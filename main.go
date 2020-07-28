package main

import (
	"datastructure"
	"fmt"
)

func main() {
	a := datastructure.Heap{}

	a.Push(31)
	a.Push(12)
	a.Push(35)
	a.Push(21)
	a.Push(5324)
	a.Push(1)

	a.PrintHeap()

	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())

}
