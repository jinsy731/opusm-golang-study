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

# Share by communicating 

동시성 프로그래밍은 광범위한 주제이기 때문에 여기에서는 Go에 한정된 중요한 내용에 대해서만 언급한다.  

__`공유변수` 에 대한 정확한 액세스를 구현하는 데 요구되는 세부 사항들로 다양한 환경에서의 동시성 프로그래밍이  
어려워졌다.__  
> "많은 환경에서 다수의 고루틴이 작동하고 있고, 그것들이 변수에 대한 액세스를 공유하는 경우에 동시성이 어렵다. "  
> ~ Todd ~ 

![img.png](img.png)  

위 그림에서, IncCounter() 라고 하는 고루틴이 두 개 존재하고, Counter 라는 공유 변수가 주어졌다.  
고루틴은 Counter 를 읽을 수 있고, 다른 고루틴도 마찬가지다.  

`스레드 양보(Yield Thread)` 라는 것은 실행 흐름을 다른 스레드에게 넘긴 다는 것으로 병렬로 작성할 수도 있다.  
고루틴들은 Counter 변수를 0부터 읽는다. 그 다음 1씩 증가 시킨다.  
각각의 고루틴이 0에서 1로 증가시킨 다음에 공유 변수에 Counter 를 다시 기록하고 있다.  

둘이서 공유 변수를 1씩 증겨 시켰다면 Counter 가 2가 되는 것이 정상적인 상황이지만, 그러지 못하고 있다.  
프로세스가 중첩하는 방식 때문에 이들이 공유 변수를 읽고 둘 다 증가시킨 다음, 둘 모두 변수를 다시 작성하고 있다.  
이는 공유 변수에 대한 읽기 및 쓰기가 잘못된 경우이다.  

이것을 `경쟁 상태(Data race)` 라고 한다.  

> Do not communicate by sharing memory. instead, share memory by communicating.  
> ~ Effective Go ~
 
경쟁상태를 해결하는 방법으로는 ___뮤텍스, 아토믹, 채널___ 이 있다.  

# Go routines  

고루틴은 단순한 함수 모델을 갖는다.  
고루틴은 같은 주소 공간에서 다른 고루틴과 동시에 실행되는 함수이다.  

고루틴은 OS의 멀티 스레드에 멀티플렉싱 된다.  
입력 및 출력 작업을 위해 대기 중일 때와 같이 하나의 고루틴이 블락되면 다른 고루틴이 계속 실행된다.  
이런 설계로 인해 여러 복잡한 생성과 관리가 드러나지 않는다.

> 멀티플렉서 (또는 MUX) 는 여러 아날로그 또는 디지털 입력 신호 중 하나를 선택하여  
> 선택된 입력을 단일 라인에 전달하는 장치이다.   
> ~ 위키백과 ~  

Go 에서 고루틴을 사용하기 위해서는 `go` 키워드를 함수나 메소드 앞에 붙인다.
하나의 고루틴에서 새 고루틴을 호출하여 실행했을 때, 하나의 고루틴 호출이 완료되면 새 고루틴은 자동으로 종료된다.  
(새 고루틴이 어떤 작업을 하고, 작업을 완료시키는 것을 기다려주지 않기 때문,  
공식 설명에서는 백그라운드에서 명령을 수행하는 유닉스 쉘 및 표기법과 유사하다고 한다. )  

또한 고루틴에 사용된 함수가 반환값을 갖는 경우, 함수가 완료될 때 삭제된다.  
(반환값을 갖는 함수를 고루틴에 사용할 때, 함수 리터럴이나 또 다른 함수로 래핑해 사용하기를 권장한다. )

`go list.Sort()`

함수 리터럴은 고루틴 호출에 유용하다.     
(익명함수를 뜻하는 것 같다. ) 

```go
func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)   
    } ()
}
```

함수 리터럴은 클로저이다.  
함수가 참조하는 변수를 사용하는 동안에는 그 생존을 보장하는 방식으로 구현되어 있다.  
다만, 위 샘플은 권장하지 않는 방식이다. 함수가 종료를 알릴 방법이 없기 때문이다.

이를 위해서는 `채널` 이 필요하다.  