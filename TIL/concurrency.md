# 동시성과 병렬성 

Go 언어로 작성된 프로그램을 오직 1개의 CPU를 가진 머신에서 실행할 경우,  
그 프로그램은 병렬로 실행되지 않을 것이다.  

동시에 실행되지 않기 때문에, 동시에 실행되는 코드의 다중 스레드가 병렬로 주어지지 않을 것이다.  
대신 순차적으로 실행된다.  

CPU가 하나이기 때문에 명령어가 하나씩 차례로 이어질 것이다.
하지만 2개 이상의 CPU를 가지고 있을 경우, 코드를 병렬로 실행할 수 있는 기회가 주어진다.  

동시성이란 무엇일까?  
동시성은 디자인 패턴이다. 코드를 작성하는 방식이다.  
싱글 코어 CPU에서 하나의 프로세스가 멀티 스레드 코드를 갖고 번갈아 가며 실행되도록 한다.  

여기까지 정리하자면, 병렬성은 물리적인 개념이고, 동시성은 논리적인 개념이다.  

> 동시성은 CPU Core가 1개일 때, 여러 프로세스를 짧은 시간동안 번갈아 가면서  
> 연산을 하게 되는 시분할 시스템으로 실행된다.  
> 병렬성은 CPU Core가 여러개일 때,  각각의 Core가 각각의 프로세스를 연산함으로써   
> 프로세스가 동시에 실행되는 것이다.  
> 출처: 인프런 개발남노씨, 기출로 대비하는 개발자 전공면접 [CS 완전정복] 

# 동시성 코드  

Go 언어에서 동시성 코드를 작성하는 방법은 간단하다.  
`go` 키워드를 이용하면 된다.  

```go
package main 

import (
	"fmt"
	"runtime" 
)

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine()) // 1
	
	go foo() 
	bar()
	
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine()) // 2
}

func foo() {
	for i:=0; i<10; i++ {
		fmt.Println("foo ", i) 
    }
}

func bar() {
	for i:=0; i<10; i++ {
		fmt.Println("bar", i)
    }
}
```

위 샘플을 실행해 보면, 처음에는 고루틴이 1개 였지만  
`go` 키워드가 사용된 `foo()` 를 만나면서 고루틴이 2개가 되는걸 확인할 수 있을 것이다.  
그런데 bar()의 내용이 출력되지만, foo()의 내용은 출력되지 않는 것도 확인될 것이다.  

처음에 고루틴이 1개 였던 이유는 main()이 이미 1개의 고루틴으로 실행되기 때문이고,  
또 다른 고루틴인 foo()를 만나면서, 이 foo()가 다른 루틴으로 처리되기 때문이다.

이때, 흐름제어는 foo()의 처리를 전혀 기다려주지 않으면서 main() 안에 명령문 처리를 계속 진행하게 된다.  
그러면 foo()의 처리가 끝나기도 전에 main()이 먼저 종료될 것이고, foo()도 같이 종료되버릴 것이다.  

따라서 일종의 동기화 작업이 필요한 상황이다.  
작성한 코드에게 기다리라는 작업을 해줄 필요가 있는 것이다.  

이 작업에는 `Mutex(상호배제 잠금)`, `WaitGroup` 등이 필요하다.  
이들은 코드가 실행을 마칠 때까지 기다려달라는 기능을 수행할 수 있다.  
여기서는 `WaitGroup` 을 이용한다.

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Added, Here! 
var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine()) // 1

	wg.Add(1) // 기다려야 할 작업(고루틴) 1개 추가 
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine()) // 2
	wg.Wait() // 작업이 끝날 때까지 기다려!
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo ", i)
	}
	wg.Done() // WaitGroup 에 등록된 작업 끝났어!  
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}
```

위 샘플을 실행하면 main() 은 종료되기 직전에 WaitGroup에 등록된 작업들이 끝날때 까지 기다려 줄 것이다.  
bar()와 고루틴의 개수가 출력되고 나서 foo()에 대한 작업 내용을 확인할 수 있을 것이다.  

