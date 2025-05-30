package main

import "fmt"

type key func(a, b int) int

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Apply(a, b int, k key) int {
	return k(a, b)
}

func main() {
	a := 5
	b := 10
	k := Multiply
	fmt.Println(Apply(a, b, k))
}
