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

`슬라이스(Slice)` 는 같은 데이터 타입을 한데 묶어준다. (=리스트)  
`Array` 는 길이가 고정되어 있는 정적 자료구조이지만, `Slice` 는 길이가 동적인 자료구조이다.  

내부적으로 Array의 첫 번째 원소를 가리키는 포인터, Slice 길이, 용량을 나타내는 값을 가진 구조이다.  
(슬라이스는 배열을 참조하고 있는 참조타입이다. )



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

### Slice Appending   
`append` 함수로 슬라이스 맨 뒤에 값을 추가할 수 있다.  
포맷은 `func append(slice []T, elements ...T) []T` 이고,   
여기서 `...T` 는 타입T의 데이터 수가 여러 개 입력 가능하다는 의미이다.  

````go 
x:=[]int{1,2,3}
x = append(x, 4, 5, 6} 
fmt.Println(x) // [ 1 2 3 4 5 6 ] 
````

또는 

```go
x:=[]int{1,2,3}
y:=[]int{4,5,6}
x = append(x, y...) 
fmt.Println(x) // [ 1 2 3 4 5 6 ]  
```

```go 
x = appned(x[:2], x[4:]...)
fmt.Println(x) // [ 1 2 4 5 6 ] 
```

위 샘플처럼 슬라이싱과 결합해 사용할 수 있다.  
이 경우에는 값을 제거하는 쪽에 가깝다.


```go
x:=make([]int, 10, 100) // make(type, len, cap) 
fmt.Println(x)      // [ 0 0 0 0 0 0 0 0 0 0 ] 
fmt.Println(len(x)) // 10
fmt.Println(cap(x)) // 100
```

`make` 는 동적으로 원소 개수가 달라지는 슬라이스에 원하는 길이와 용량을 정할 수 있도록 해준다.  
(용량에 대한 매개변수 생략 가능하다. ) 
`make` 로 생성된 빈 슬라이스는 원소를 추가하면 설정한 용량만큼 길이가 증가할 것이다.  
그런데 여기서 슬라이스에 저장된 원소의 개수가 100으로 지정된 용량을 넘기면 어떻게 될까?  
슬라이스는 자신의 용량(정확히는 슬라이스가 가리키는 하부 배열 용량)을 두 배로 늘려 기존의 원소를 잃지 않고 계속 사용할 수 있도록 한다.
정확히는 기존에 가리키던 하부 배열의 원소를 그대로 들고있으면서, 용량이 두 배인 배열을 새로 만들고, 이전 배열과 새 배열을 가리키고 있는 것이다.

### Multiple Slice 

```go 
package main 

import "fmt" 

func main() {
    first := {"Kim", "Lee", "Jeon", "Kwon"} 
    name := {"Hank", "Mingi", "Wardy", "Hansung"}
    clasi := [][]stirng{first, name} 
    
    fmt.Println(clasi)
}
```

## Map

`map` 은 `key:value` 쌍으로 이루어진 자료구조이다.  

```go
package main 

import "fmt"

func main() {
	m := map[string]int{ 
		"Bond" : 32,
		"Penny" : 27,
    }
	fmt.Println(m)
	v, ok := m["Donkey"]
	fmt.Println(v, ok) 
}
```

위 샘플 코드를 보듯이, `map[<key type>]<value type>` 포맷으로 정의하고, `{ }` 안에  필요한 값을 입력한다.  
`v, ok := m["Donkey"]` 는 맵에 Donkey란 키를 넣었을 때 값이 있는지 없는지 체크하는 구문이다.  
이를 Go에서는 관용적으로 조건문과 함께 쓰는데 그 모습은 아래 샘플과 같다. (`comma ok` 라고 함. ) 

```go 
if v, ok := m["Donkey"]; ok {
    fmt.Println(v)  
} 
```


`for` 반복문에서도 `range` 와 함께 사용할 수 있다.  

```go
for k, v := range m { 
	fmt.Println(k, v)
}
```

`map` 의 값을 삭제할 수도 있다.  

```go
delete(m, "Bond") 
```

다만 `map` 을 삭제할 때는 없는 키값을 입력해도 에러가 안나기 때문에  
`if v, ok := <map name>["map key"]; ok { ... }` 로직으르 처리하길 권장한다.  

```go
package main 

import "fmt" 

func main() {
	m := map[string]string {
		"OpusM" : "LedgerMaster",
		"Hexlant" : "Octet", 
		"Pala" : "PalaSquare",
    }
	
	if v, ok := m["Krust"]; ok {
		fmt.Println(v) 
		delete(m, "Krust")
    }
}
```

## Struct

`struct` 는 서로 다른 유형의 데이터를 함께 구성할 수 있는 데이터 구조이다.

```go
package main 

import "fmt" 

type person struct { 
	first string
	last string 
}

func main() {
	p1 := person{
		first: "James", 
		last: "Bond",
    }
	
	p2 := person{ 
		first: "Miss",
		last: "Moneypenny",
    }
	
	fmt.Println(p1) // {James Bond} 
	fmt.Println(p2.fisrt, p2.last) // Miss Moneypenny 
}
```

정의한 `struct` 의 값을 가져올 때는 `점 표기법(Dot notation)` 을 사용한다. 

## Embedded Struct

한 타입을 선택하고 그것을 다른 타입에 임베드하는 방법이다. 

```go
package main

type person struct {
	first stirng
	last string
}

type secretAgent struct {
	person 
	ltk bool 
}

func main() {
    //p1 := person{ 
	//	first: "James",
	//	last: "Bond",
    //}
	
	sa1 := secretAgent{
		person: person{ 
			first: "James",
			last: "Bond",
        },
		ltk: true,
    }
	
	fmt.Pritnln(sa1.first, sa1.last, sa1.ltk)
}
```

위 샘플에서는 `person` 이 임베디드 구조체이며, `secretAgent` 에 임베드 되었다.  
임베드된 구조체를 `inner type` 이라 하고,  
내부 필드가 승격되어 상위 타입(`outer type`)에서 내부필드에 바로 접근할 수 있다.  

## Anonymous Struct

구조체를 `익명 구조체` 로 만들어 표현할 수도 있는데,

```go
pacakge main 

import "fmt" 

func main() {
    p1 := struct {
        first string
        last string 
    } { 
        first: "James",
        last: "Bond", 
    }   
}
```

라고 코드를 작성할 수 있다. 이러한 방식은 코드를 깨끗하고 오염없이 린(lean) 하게 유지하고 싶을 때 사용한다.  
필요 없는 외부 타입이나 변수를 사용하지 않는다면 이런 식으로 작성하는 것을 권장한다.
