package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func QuizThree() string {
	s := fmt.Sprintf("%d %s %t\n", x, y, z)
	// s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	return s
}

func main() {
	fmt.Print(QuizThree())
	// fmt.Println(QuizThree())
}
