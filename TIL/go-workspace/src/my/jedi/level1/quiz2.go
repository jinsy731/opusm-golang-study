package main

import "fmt"

var x int
var y string
var z bool

func QuizTwo() {
	// Zero values
	fmt.Printf("x: %d\n", x)
	fmt.Printf("y: %s\n", y)
	fmt.Printf("z: %t\n", z)
}

func main() {
	QuizTwo()
}
