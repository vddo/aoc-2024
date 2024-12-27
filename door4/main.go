package main

import (
	"door4/arraystack"
	"fmt"
)

func main() {
	fmt.Println("Program door4 starting...")

	a := arraystack.New()
	a.Push('h')
	a.Push('i')
	fmt.Println(a)
	fmt.Println(a.Pop())
	fmt.Println(a)

	fmt.Println(a.Empty())
	fmt.Println(a.Pop())

	fmt.Println(a.Empty())
}
