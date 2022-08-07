# Go Method Sets 

메소드 세트는 타입에 첨부되는 메소드를 결정한다.  
메소드 세트에는 두 종류가 있는데,

* 논 포인터 리시버 (`Non pointer receiver`)
* 포인터 리시버 (`Pointer receiver`) 

위 두 종류이다.  

```text
Receivers    Values    
--------------------
(t T)        T and *T
(t *T)       *T 
```

## 논 포인터 리시버  

포인터가 아닌 리시버, 리시버가 있는 타입에 함수를 첨부한다.  
포인터 또는 포인터가 아닌 값과 함께 작동한다.  

```go
package main 

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64 
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius 
}

func info(s shape) { 
	fmt.Println("area", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}
	info(c)
}
```

## 포인터 리시버  

오직 포인터인 값에서만 작동한다.

```go
package main 

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64 
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius 
}

func info(s shape) { 
	fmt.Println("area", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}
	// fmt.Printf("%T\n", &c) // *main.circle 
	info(&c)
}
```

