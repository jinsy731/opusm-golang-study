# Pointer

`포인터(Pointer)` 는 값이 저장된 메모리의 주소를 가리키는 것이다.  
우편함(주소) 안에 편지(값)가 있는 것처럼 말이다. 간단하다.

```go
package main 

import "fmt" 

func main() {
	a := 123 
	fmt.Println(&a) // 0x1040a124
	fmt.Printf("%T\n", &a) // *int
}
```

`Pointer` 에서 `&` 기호를 변수명 앞에 붙여줘서 해당 변수의 메모리 주소를 확인할 수 있다.  
위 샘플에서 `&a` 는 메모리 주소 `0x1040a124` 를 담고 있고, 타입을 확인해보면  
`*int` 이다. `pointer int` 라고 하고, `int` 에 대한 포인터라는 의미이다.  
`int` 와 `*int` 는 엄연히 다른 타입이다. 따라서, `*int` 를 어떤 변수에 할당하려면 아래와 같이 해야 한다.   

```go
var pa *int = &a 
```

그리고 할당된 주소에 대한 값을 얻으려면 `역참조(Dereference)` 를 해야하고, 방법은 아래와 같다.  

```go
fmt.Println(*pa) // 123
fmt.Println(*&a) // 123
```

포인터는 하나의 값을 가리키는 메모리 주소를 갖고 있기 때문에, 포인터가 가리키는 값을 변경하면 그 값을 참조하고 있던  
다른 변수들의 값도 모두 변경된다.  

```go
package main 

import "fmt" 

func main() {
	a := 123
	b := &a 
	
	*b = 234
	fmt.Println(*b) // 234
	fmt.Println(a) // 234 
}
```

참고로 Go언어는 모든게 ___Pass by value___ 인 점을 명심하자. 

# When use

1. 용량이 큰 데이터를 전달해야 할 때
2. 특정 위치의 데이터를 변경해야할 때(`mutate`)