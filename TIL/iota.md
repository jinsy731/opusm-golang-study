# iota
Golang에 정의된 특수한 식별자 `iota` 는 0부터 1씩 증가시켜 주는 특수한 역할을 한다.  

아래 샘플을 보자. 

```go
package main

import "fmt" 

const ( 
	_ = iota // 0부터 시작하므로 무시
	KB  = 1 << (10 * iota)         // 2^10
	MB                             // 2^20
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	// ... 
}
```

위 샘플에서 iota는 0부터 시작하므로 `_` 으로 첫 iota를 무시하고,  
다음 iota인 1을 shift해 2^10인 값을 만들어 주고 있다.  
그 후 iota는 정의된 const 값들 만큼 1씩 증가해 대입된다.  

```
[ M E M O ]
* iota를 bit shifting에서 사용하는 것을 golang spec에서  
    권장하고 있다.  
    일반적인 개발자가 보기엔 가독성이 떨어지지만, 
    golang spec에서는 코드를 더 간결하고 스마트하게 사용하길 원한다.  
```