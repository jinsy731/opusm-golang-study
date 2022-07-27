# Go Loop 

Go Loop는 우리가 익히 알고있는 반복분을 의미하고,  
여느 언어들 처럼 `for <init> ; <condition>; <post> { .. }` 로 구조화된다.   

아래는 `for` 를 이용한 반복문 샘플이다.    
샘플을 보며 go loop를 어떻게 사용할 수 있는지 파악하자. 

```go
package main 

import "fmt" 

func main() {
    i := 1
    for i<=3 {
        fmt.Println(i)
        i = i + 1
    }
	
    for j:=7; j<=9; j++ {
       fmt.Println(j)
    }  
	
    for {
        fmt.Println("loop") 
		break
    }
	
	for n:=0;n<=5;n++ {
		if n%2 == 0 {
			continue  
        }
		fmt.Println(n)
    }
}
```

출처: https://gobyexample.com/for