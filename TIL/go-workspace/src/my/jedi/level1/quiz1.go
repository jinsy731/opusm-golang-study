package main

import "fmt"

func QuizOne() {
	x := 42
	y := "James Bond"
	z := true

	// multiple statements
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// single statement
	fmt.Printf("%d, %s, %t\n", x, y, z)
}

func main() {
	QuizOne()
}