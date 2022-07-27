# Go 자료구조

들어가기 앞서 `Composite data type`, `Aggregate data type`, `composite literal` 에 대해서 배워보자.  

### Composite Data Type

`Composite Data Type` 이란 원시 데이터 타입과 다른 다양한 데이터 타입의 값들이 있는  
`합성 데이터 타입` 이다.  
이렇게 구조화한 데이터 타입을 `구조체(struct)` 라고 한다.  
`struct` 는 이 글의 아래에 정리되어 있으니 여유를 갖고 읽기 바란다.

### Aggregate Data Type

참고로 `Aggregate Data Type` 이란 것도 있는데 `집합 데이터 타입` 이라 하고,  
`Array` 나 `List` 가 이 포지션에 속한다.

### Composite Literal

불필요한 코드를 줄이고 한데 모아 새로운 인스턴스를 생성한 것  
`합성 리터럴(Composite Literal)` 은 배열, 슬라이스, 맵을 생성하는데 사용될 수 있다.  
합성 리터럴의 필드들은 순서대로 배열되어 있어야 한다.  
필드 레이블로 인덱스와 멥의 키를 적절히 사용해야 한다.

합성 리터럴은 아래와 같은 형태로 정의될 수 있다.

```go
a := [...]string{"no err", "Eio", "in"}
s := []string{"no err", "Eio", "in"}
m := map[string]interface{}{"name": "hihi", "age": 24}
```

Effective Go에 소개된 참고하기 좋은 예제는 아래와 같다.

```go
// Worst practice.  
func NewFile(fd int, name string) *File {
  if fd < 0 {
    return nil
  }
  f := new(File)
  f.fd = fd
  f.name = name
  f.dirinfo = nil
  f.nepipe = 0
  return f
}
```

```go
// Best practice.
func NewFile(fd int, name string) *File {
  if fd < 0 {
    return nil
  }
  // f := File{fd, name, nil, 0}
  return &File{fd: fd, name: name} 
}
```

위의 worst practice는 불필요한 코드가 있음에 분명하다.  
불필요한 코드는 합성 리터럴로 해결할 수 있다고 배웠고, 아래 best practice 가 좋은 방안이 될 수 있다.

합성 리터럴은 필드 레이블이 순서에 맞게 들어가야 하며, 이는 값에 `라벨링`을 해줌으로 해결한다.  
따라서 `File{fd, name, nil, 0}`는 `File{fd: fd, name: name}` 으로 합성 리터럴을 이용해 코드를 작성한다.  
이 때, 3번째~4번째 자리의 값은 `Zero values` 가 들어가 생략 가능해진다.  

---
## Array

```go
package main

import "fmt"

func main() {
    var x [5]int
    fmt.Println(x) // [ 0 0 0 0 0 ]
    x[3] = 42
    fmt.Println(x) // [ 0 0 0 42 0 ]
    fmt.Println(len(x)) // 5
}
```

Zero based index를 따르는 일반적인 배열이다.  
메모리 레이아웃을 상세히 계획하는데 유용하고,  
때로는 과도한 메모리 할당을 피하는데 도움이 된다.  
아쉽지만, Go에서는 자주 사용되지 않는다.  
Effective Go 문서에서는 Array 대신 Slice를 이용하길 권장한다.  
(Go는 쉬운 프로그래밍, 효율적인 실행과 컴파일을 지향함을 잊지말자. )

```go
[ M E M O ] 
* 아래 사이트를 자주 애용하자. 책보다 더 값진 무료 레퍼런스이다.  
* https://golang.org/doc/effective_go.html 
```

```text
[ P O I N T ]
[5]int와 [6]int는 같은 타입일까?
답은 `No` 이다.  
왜냐하면 Array의 길이도 타입의 일부이기 때문이다.
```

## Slice

```go
package main 

import "fmt" 

func main() {
    // x := type{values} -- composite literal
    x := []int{4, 5, 6, 7, 8} 
}
```

`슬라이스(Slice)` 는 같은 데이터 타입을 한데 묶어준다.  

### Slicing 

슬라이스를 슬라이싱하는 방법이다.  
슬라이싱은 `:` 기호를 이용한다.  

```go
package main 

import "fmt" 

func main() {
    x := []int{1,2,3,4,5} 
    fmt.Println(x[1])       // 5
    fmt.Println(x)          // [ 1 2 3 4 5 ]
    fmt.Println(x[1:])      // [ 2 3 4 5 ]
    fmt.Println(x[1:3])     // [ 2 3 ] 
} 
```

위 샘플에서도 보듯이, 슬라이스를 슬라이싱하는 `:` 기호는  
`x[n:n-1]` 과 같이 사용한다.  
슬라이싱은 n부터 시작해서 n-1까지를 잘라 가져온다는 의미이다.  

## Map

## Struct

## House Keeping 