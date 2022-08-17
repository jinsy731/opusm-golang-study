# Testing

테스트 코드를 작성하는 방법을 배운다.  
테스트를 하는 이유는 오류 처리에 더 최적화된 프로그래밍을 할 수 있기 때문이다.  

테스트를 만들려면 파일명 끝에 `_test.go` 를 꼭 붙여줘야 한다.  
그리고 테스트를 원하는 파일과 동일한 패키지 내에 있어야 하고,  
함수명은 모두 `func TestXXX(t *testing.T)` 이어야 한다.  

정리하자면  
* 파일 이름은 파일명_test.go
* 같은 패키지에 위치  
* func TestXXX(t *testing.T)

[ main.go ]  
```go
package main 

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", mySum(2,3))
	fmt.Println("4 + 5 = ", mySum(4,5))
	fmt.Println("12 + 13 = ", mySum(12,13))
}

func mySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
    }
	return sum 
}
```

[ main_test.go ]  
```go
package main 

import "testing" 

func TestMySum(t *testing.T) {
	x := mySum(2, 3) 
	if x != 5 {
		t.Error("Expected", 5, "Got" , x) 
    }
}
```

[ command ]  
`go test -v`  

# 테이블 테스트  

테이블 테스트는 쉽게 말하면 테스트의 시리즈라고 볼 수 있는데,  
데이터들이 여러 개 있고 그 데이터들이 특정 함수에서 잘 동작하는지  
다중 테스트를 통해서 검사하는 방식이다.  

[ main_test.go ]  
```go
package main 

import "testing" 

func TestMySum(t *testing.T) {
	
	type test struct { 
		data []int 
		answer int 
    }
	
	tests := []test{
		test {
			[]int{21,21},
			42,
		},
		test {
			[]int{3,4,5},
			12,
		},
		test {
			[]int{1,1},
			2,
		},
		test {
			[]int{-1,0,1},
			0,
		},
    }
	
	for _, v := range tests {
		x := mySum(v.data...) // 데이터 펄링 
		if x != v.answer {
			t.Error("Expected", v.asnwer, "Got", x) 
        }
    }
}
```

# 벤치마크 테스트

[ main_test.go ]
```go
package main 

import "testing" 

func BenchmarkMySum(b *testing.B) {
	
	for i := 0; i < b.N; i++ {
		mySum(2, 3)
    }
	
}
```

[ command ]   
`go test -bench <identifier>` // go test -bench MySum    
참고로, 전체 벤치마크 테스트는 `go test -bench .`   

# 커버리지

커버리지란 프로그래밍에서 작성한 코드가 테스트에 얼마나 반영되는지 측정하는 역할이다.  
즉, 테스트에 적용된 코드가 많을수록 `Best Practice(모범 사례)` 가 된다.  
하지만 현실적으로 코드 전체가 100% 테스트에 반영되긴 어렵기 때문에  
70~80% 반영되었다면 "잘 했다. " 라고 할 수 있다.  

커버리지는 테스트 코드가 작성된 패키지로 가서 실행한다.  
실행 명령어는 다음과 같다.  

`go test -cover`  

뮈 명령어를 실행하면 테스트와 함께 커버리지를 보여줄 것이다.  
```text
PASS
coverage: 100.0% of statements 
```

파일로 내보내는 방법도 있는데 수치는 보여주지 않을 것이다.   
`go test -coverprofile c.out`  

html로 보여주는 방법도 있다.  
이 방법은 어떤 코드가 커버리지 되었는지 안되었는지 확인할 수 있어서 좋다.  
`go tool cover -html=c.out`  


