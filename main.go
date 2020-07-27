package main

import (
	"datastructure"
)

func main() {
	a := datastructure.NewTree()

	a.Insert(13)
	a.Insert(5)
	a.Insert(162)
	a.Insert(234)
	a.Insert(6)
	a.Insert(8)
	a.Insert(6)
	a.Insert(162)
	a.Insert(162)
	a.Insert(2)
	a.Insert(1)
	a.Insert(150)
	a.Insert(132)

	a.RecurSiveDFS()
	a.StackDFS()

	a.BFS()
}
