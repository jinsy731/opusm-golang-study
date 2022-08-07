# Json Marshal

JSON은 Javascript Object Notation의 약자이고,   
Key-Value Pair로 이루어진 배열 자료형 또는 데이터 오브젝트를 표현한다.  

현재 인터넷에서 데이터를 주고 받을 때 주로 사용되는 자료구조 중 하나이다.  

Json Marshal 이란, Go언어에서 데이터를 Json 포맷으로 Encoding 해주는 것이고,  
`encoding/json` 패키지의 `Marshal()` 함수를 사용한다.  

```go
package main 

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last string 
	Age int 
}

func main() {
    p := person {
		First: "James",
		Last: "Bond",
		Age: 42,
    }	
	
	p2 := person {
		First: "Miss",
		Last: "Moneypenny",
		Age: 27, 
    }
	
	people := []person{p, p2} 
	
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error: ", err) 
    }
	
	fmt.Printf("%T\n", bs)
	fmt.Println(bs)
}
```

# Json Unmarshal 

Json Unmarshal은 Json 포맷 자료구조를 Go 데이터로 변환하는 것이다.  
Easy!  

```text
// (1) 
// Field appears in JSON as key "myName".
Field int `json:"myName"`

// (2)
// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// (3) 
// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// (4) 
// Field is ignored by this package.
Field int `json:"-"`

// (5) 
// Field appears in JSON as key "-".
Field int `json:"-,"`
```

우선 구조체 필드 태그(struct field tags)에 대해서 알아보자.  
위 샘플은 http://go.dev/pkg/encoding/json 에서 가져온 것이고, 설명을 위해 번호를 추가했다.  

(1) Field를 Json에서는 myName 이라는 Key로 나타낸다는 의미이다.  
(2) Field를 Json에서 myName 이라는 Key로 나타내고, 비어있을 때 객체에서 생략한다. (`omitempty`)  
(3) Field를 Json에서 Field 그대로 나타내고(default), 비어있으면 객체에서 생략한다.  
(4) Field를 패키지에서 무시한다. 이와 관련한 작업은 아무것도 하지 않는다.  
(5) Field를 Json에서 `-` 기호로 나타낸다.  

다음으로 Json Unmarshal 해보는 샘플을 확인해보자.  

```go
package main 

import (
	"encoding/json"
	"fmt" 
)

type person struct {
	First string
	Last string 
	Age int 
}

func main() {
	// (1) Json format data define 
	data := `[{"First":"James", "Last":"Bond", "Age":32}, {"First":"Miss", "Last":"Moneypenny", "Age":27}]`
	// (2) Convert data to bytes slice
	bs := []byte(data) 
	// (3) Define data for unmarshal 
	// people := []person{} 
	var people []person 
	// (4) Try unmarshal and print error 
	err := json.Unmarshal(bs, &people) 
	if err != nil {
		fmt.Println("error: ", err) 
    }
	// (5) Print data 
	for i, v := range people {
		fmt.Println("Index: ", i, " : ", v.First, v.Last, v.Age)
    }
}
```

# References

* https://mholit.github.io/json-to-go  
  * Json 포맷 데이터를 Go 데이터로 변환해주는 사이트  
