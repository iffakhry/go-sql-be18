package main

import "fmt"

func main() {
	defer fmt.Println("hello 1")
	defer fmt.Println("hello 2")
	functionName()
	fmt.Println("hello 3")
	fmt.Println("hello 4")
}

func functionName() {
	defer fmt.Println("func 1")
	panic("error")
	fmt.Println("func 2")
}
