package main

import (
	"datastructure"
	"fmt"
)

func main() {

	m := datastructure.Map{}
	m.Add("kangseungho", "01044982128")

	fmt.Println(m.Get("kangseungho"))

}
