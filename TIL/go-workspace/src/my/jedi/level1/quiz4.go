package main

import "fmt"

type MyInt int

var x MyInt

func QuizFour() {
    fmt.Println(x)
    fmt.Printf("%T\n",x)
    x = 42
//     x = int(42)
    fmt.Println(x)
}

func main() {
    QuizFour()
}