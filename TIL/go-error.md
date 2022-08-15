# Error handling

왜 Go에는 예외처리가 없을까?

Go 를 개발한 사람들은 
> try, catch 형태로 예외를 처리하는 구조가 코드를 난해하기 만든다.  
> 이렇게 처리하면 간단한 에러도 개발자들은 다 예외로 처리해버리고 만다.    

라는 생각을 가지게 되어 Go 만의 에러처리 매커니즘을 만들게 되었다.  

아래 참고자료를 꼭 읽어보자.  

* Go에 예외처리가 없는 이유  
  * https://golang.org/doc/faq#exception
* Go에서 에러처리 방법
  * https://blog.golang.org/error-handling-and-go
* Error package 
  * https://godoc.org/errors

프린팅과 로깅하는 방법에 대해서도 알아보자.
* fmt.Println()
    * log.Println() // 그냥 로그를 남겨야할 상황
* log.Fatalln() // 엄청나게 치명적인 상황
    * os.Exit() // 로그 메시지를 출력 후, os.Exit(1) 명령을 호출
* log.Panicln() // 치명적인 상황과 치명적이지 않은 상황 중간 단계
    * deferred functions run
    * can use "recover"
* panic() // __빌트인 패키지에 있고, 현재 실행 중인 고루틴의 실행을 정지, defer는 실행됨__

```go
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("no-file.txt")
	
	defer file.Close() // 지연 기능, 마지막 명령까지 실행하고 나서 뒤늦게 처리됨 
	
	if err != nil {
		//fmt.Println("error happend: ", err)
		log.Println("error happend: ", err)
		//log.Fatalln(err)
		//panic(err) 
	}
}
```

위 샘플은 fmt, log, panic 을 사용해보고 있다. 

## fmt와 log의 차이  

log는 타임스탬프를 남겨준다.
(끝) 

### log 파일 기록하기 

```go
package main 

import (
	"log"
	"fmt" 
	"os"
)

func main() {
	f, err := os.Create("log.txt") 
	if err != nil {
		log.Println("Error: ", err) 
    }
	defer f.Close()
	log.SetOutput(f)
	
	fmt.Println("----exit----")
}
```

## 에러를 회복하는 방법, Recover 

`Recover` 는 에러에서 회복하는 방법이다.  
빌트인 패키지에 포함되어 있고, 패닉 상태에 대해서 제어할 수 있다.  

리커버는 defer 된 함수 안에서만 쓸모 있다.  
보통 상태에서 실행 중에 리커버를 호출하면 nil 이 반횐되며 다른 효과는 없다.  

현재 실행 중인 고루틴이 패닉 상태에 빠져 리커버를 호출하면 패닉에 주어진 값이 저장되고,  
일반 실행을 재개한다. 

```go
package main

import "fmt"

func main() {
	f() 
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r:=recover(); r != nil {
			fmt.Println("Recoverd in f ", r)
        }
    }()
	fmt.Println("Calling g.")
	g(0) 
	fmt.Println("Returned normally form g.")
}

func g(i int) {
	if 3 < i {
		fmt.Println("Panic! ")
		panic(fmt.Sprintf("%v", i))
    }
	defer fmt.Println("Defer in g ", i)
	fmt.Println("Printing in g ", i)
	g(i+1) 
}
```

위 샘플에서 i 가 4가 되어 패닉이 발생한다.  
패닉이 발생하면 콜스택을 후입선출로 거슬러 올라가 3, 2, 1, 0 올라가 defer 들이 실행되고,  
리커버가 패닉에 저장된 값을 실행한다.  

## 에러를 생성하는 방법

```go
package main 

import (
	"fmt"
	"errors"
)

//var ErrNorgateMath = errors.New("norgate math: square root of negative number.")

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err) 
    }
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, ErrNorgateMath
		return 0, erros.New("norgate math: square root of negative number.")
    }
	return 42, nil 
}
```

위 방식 처럼 성공에 대한 반환은 nil 로 에러가 없음을 반환하고,  
어떤 조건이 성립해 그것이 에러라면 `errors.New()` 로 에러를 생성하고 반환할 수 있다.  

```go
package main 

import (
	"fmt"
	"errors"
)

//var ErrNorgateMath = errors.New("norgate math: square root of negative number.")

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err) 
    }
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, ErrNorgateMath
		return 0, fmt.Errorf("norgate math again: squre root of negative number: %v", f)
    }
	return 42, nil 
}
```

위와 같이, fmt.Errorf() 를 이용할 수도 있다.  
참고로 fmt.Errorf() 는 내부적으로 errors.New() 로 구현되어 있다.  

```go
package main 

import (
	"fmt"
)

type norgateMathError struct {
	lat string
	long string 
	err error 
}

func (n norgateMathError) Error() string {
	return fmt.Sprinf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23) 
	if err != nil {
		log.Println(err)
    }
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errof("norgate math redux: square root of negative number")
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
    }
	return 42, nil
}
```

# References

* Defer, Panic, and Recover
  * https://blog.golnag.org/defer-panic-and-recover  
