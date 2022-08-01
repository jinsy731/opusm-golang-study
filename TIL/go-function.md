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
}
```

가변 매개변수를 활용하는 샘플은 위의 코드와 같다.  
이 때, 가변 매개변수는 `<name> ...<type>` 으로 표현하는데  
타입을 출력해보면 알 수 있듯이, 가변 매개변수의 경우에는 타입이 `slice` 이다.   
