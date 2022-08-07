# Go Sort 

슬라이스의 값들을 정렬한다.  

```go
package main 

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}
	
	sort.Ints(xi) 
	sort.Strings(xs) 
	
	fmt.Println(xi) 
	fmt.Println(xs) 
}
```

# Sort Custom

```go
package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

type ByAge []person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	p1 := person{
		"James",
		32,
	}
	p2 := person{
		"MoneyPennry",
		27,
	}
	p3 := person{
		"Q",
		64,
	}

	people := []person{p1, p2, p3}
	var bA ByAge = ByAge(people)
	sort.Sort(bA)
	fmt.Println(people)
}
```

ByAge 를 대상으로 Sort를 구현하기 위해서는 Len(), Less(), Swap() 을 구현해야 한다.  
왜냐하면 위 함수들이 sort.Sort 에 정의된 Interface 가 필요하기 때문인데  
이 Interface는 다시 Len(), Less(), Swap()을 구현할 것을 요청하고 있다.  

```go
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

따라서 ByAge 는 []person 에 대한 sort.Interface 를 구현한다.  

# References

* https://pkg.go.dev/sort@go1.19
* https://godoc.org/sort