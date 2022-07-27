# range 

문자열, 배열, 맵 등의 자료구조를 다룰 때 `range` 는 해당 자료구조에 담겨있는  
원소를 순회할 수 있도록 해주는 이터레이터다.  
그래서 `for` 반복문에서 잘 활용될 수 있는 유용한 친구라고 볼 수 있다.  

아래 샘플을 통해 `range` 가 `for` 문에서 어떻게 사용되는지 확인하자.  

```go
package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```

출처: https://gobyexample.com/range  
