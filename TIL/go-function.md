# Go의 함수

Go에서 함수는 아래 스니펫과 같이 표현할 수 있다. 

```go
func (r receiver) identifier(parameters) (returns) { ... } 
```

```text
[ M E M O ]
* paramters는 함수를 정의할 때 쓰는 용어 
* arguments는 함수를 호출할 때 쓰는 용어 
``` 

# 가변 매개변수  
  
```go
package main 
import "fmt"

func main() {
	testFunction(1, 2, 3)
}

func testFunction(x ...int) { 
    fmt.Println(x) // [ 1 2 3 ]
	fmt.Printf("%T\n", x) // []int
	
	sum := 0
	for _, v := range x {
		sum += v
    }
	fmt.Println(sum)
}
```

가변 매개변수를 활용하는 샘플은 위의 코드와 같다.  
이 때, 가변 매개변수는 `<name> ...<type>` 으로 표현하는데  
타입을 출력해보면 알 수 있듯이, 가변 매개변수의 경우에는 타입이 `slice` 이다.   

# 슬라이스 언펄링

언펄링의 사전적 의미가 `펼치다` `흐트러지다` 인 것 처럼, `슬라이스 언펄링` 이란 슬라이스 값을 펼치는 형태로 인자를 전달한다는 것이다.  

```go
package main

import "fmt" 

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8} // 정적 프로그램된 슬라이스
	result := allSum(xi...) // 슬라이스 언펄링, 펼쳐서 모든 요소를 전달한다는 의미
	fmt.Println(result)
}

func allSum(x ...int) int { 
	sum := 0
	for _, v := range x {
		sum += v
    }
	return sum 
}
```

💡 참고로 가변 매개변수로  아무것도 전달하지 않을 수 있다. (passing in zero or more values)

 # Defer  

호출이 끝나는 지점까지 함수 실행을 지연시킨다.  

```go
package main 

import "fmt" 

func main() {
	defer foo() 
	bar() 
}

func foo() {
	fmt.Println("I am foo") 
}

func bar() {
	fmt.Println("I am bar")
}
```

위 샘플에서 foo()와 bar() 중에 어떤 것이 먼저 호출될까?   
bar()가 먼저 호출되어 실행되고, 끝나면 foo()가 실행된다. 

# 리시버 매개변수  

함수는 `func (r receiver) identifier(parameter(s)) return(s)` 로 구성된다.  
이 중에서 `r recevier` 에 해당하는 부분을 리시버 매개변수라고 하는데,  
이 매개변수가 정의된 함수는 리시버 매개변수의 타입에 해당하는 메소드가 된다.  

```go
package main

import "fmt" 

type person struct { 
	fisrt string
	last string 
}

type secretAgent struct {
	person 
	patner string 
	ltk bool 
}

func (s secretAgent) speak() { 
	fmt.Println("I am ", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person { 
			"James", 
			"Bond",
        },
		patner: "Moneypenny",
		ltk: true,
    }
	fmt.Println(sa1)
	sa1.speak()         // 마치 메소드처럼 호출한다. 
}
```