package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	//quit := make(chan int)
	quit := make(chan bool)

	// send
	go send(eve, odd, quit)

	// recv
	recv(eve, odd, quit)

	fmt.Println("----exit----")
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func recv(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("from the even channel: ", v)
		case v := <-odd:
			fmt.Println("from  the oven channel: ", v)
		case i, ok := <-quit: // quit 채널은 닫혔을 때 값을 전달한다.
			if !ok {
				fmt.Println("not ok ", i, ok)
				return
			} else {
				fmt.Println("ok", i)
			}
		}
	}
}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	eve := make(chan int)
//	odd := make(chan int)
//	quit := make(chan int)
//
//	// send
//	go send(eve, odd, quit)
//
//	// recv
//	recv(eve, odd, quit)
//
//	fmt.Println("----exit----")
//}
//
//func recv(e, o, q <-chan int) {
//	for {
//		select {
//		case v := <-e:
//			fmt.Println("from the even channel: ", v)
//		case v := <-o:
//			fmt.Println("from the odd channel: ", v)
//		case v := <-q:
//			fmt.Println("from the quit chaanel: ", v)
//			return
//		}
//	}
//}
//
//func send(e, o, q chan<- int) {
//	for i := 1; i < 100; i++ {
//		if i%2 == 0 {
//			e <- i
//		} else {
//			o <- i
//		}
//	}
//	close(q)
//}

//package main
//
//import (
//	"fmt"
//)
//
//// main
//func main() {
//	c := make(chan int) // 양방향 채널
//
//	// send
//	go sender(c)
//
//	// receive
//	receiver(c)
//
//	fmt.Println("----exit----")
//}
//
//// send
//func sender(c chan<- int) { // 송신 전용 채널
//	c <- 42
//}
//
//// receive
//func receiver(c <-chan int) { // 수신 전용 채널
//	fmt.Println(<-c)
//}
