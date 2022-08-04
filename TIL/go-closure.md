# 클로저 (Closure) 

`클로저(closure)` 는 무언가를 닫는다는 개념이다.  
프로그래밍에서는 변수 범위가 코드의 한 영역으로 제한되도록 변수와 일부 코드를 닫는다.

```go
package main 

import "fmt" 

func main() {
	x := 10 
	
	{
		y := 20
		fmt.Println(y) // 20 
    }
	// fmt.Println(y) // Error 
	
	fmt.Println(x)
}
```

위 샘플에서 보듯이, Go 언어에서 클로저는 _중괄호(`{}`) 안에 중괄호(`{}`)_ 로 표현한다.  
`y` 같은 변수들은 자기가 속한 중괄호 내에서만 인식되기 때문에 외부에선 접근할 수 없도록 제한되었다.  

```go
package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())    // 1
	fmt.Println(a())    // 2
	fmt.Println(a())    // 3
	fmt.Println(a())    // 4
	fmt.Println(b())    // 1
	fmt.Println(b())    // 2 
}

func incrementor() func() int {
	var x int   // zero value
	return func() int {
		x++
		return x
    }
}
```

위 샘플은 이전 샘플보다 조금 더 복잡하다. 함수를 반환하는 함수가 있고, 변수 두 개에 할당되는데  
할당된 메모리 주소가 다르기 때문에 두 변수의 결과는 서로 다름을 확인할 수 있다.  

`incrementor()` 에 존재하는 변수 `x` 는 클로저가 `func() int { ... }` 까지인 것도 확인 할 수 있다.  

