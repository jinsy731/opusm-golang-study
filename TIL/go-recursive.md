# 재귀(Recursive)  

재귀는 여러모로 좋지 않다.  
불필요하게 복잡하고, 메모리에 부정적인 영향을 준다.  

그러니 그냥 재귀가 이런 녀석이구나 하는 것만 알고 넘어가면 좋겠다.  
_재귀란 함수가 자기 자신을 호출하는 것이다._  
재귀가 무엇인지는 이걸로  충분하다.  

```go
package main 

import "fmt"

func main() {
	f := factorial(4) 
	fmt.Println("Result: ", f) 
}

func factorial(n int) int {
	if(n == 0) {
		return 1
    }
	return n * factorial(n-1) // 4 * 3 * 2 * 1 * 1  
} 
```

위 샘플에서, 팩토리얼 함수는 메모리에 복귀주소와 함께 스택프레임이 차곡차곡 쌓일 것이다.  
마지막 연산이 끝날때 까지는 아마 메모리에서 제거되지 않을 것이다.   
연산이 끝나야 비로소 위에서 부터 아래로 제거되는 형태로 보일 것이다.

그런데 문제는 이런 연산이 끝나지 않는 채로 계속해서 스택프레임이 쌓이게 된다면 어떨까?  
많은 메모리 공간을 차지할 것이고, 결국 스택 오버플로를 발생시켜 프로그램은 비정상적 종료 처리가 될 것이다.  

이처럼 재귀는 안전하지 않은 지름길과 같다.  
그러므로 재귀 보다는 루프로 처리하기를 바란다. 재귀로 할 수 있는 일을 루프로도 할 수 있다. 😃

```go
package main 

import "fmt"

func main () {
	f := factorial(4)
	fmt.Println(f) 
	
	lf := lFactorial(4) 
	fmt.Println(lf)
}

func factorial(x int) int {
	if (x == 0) { return 1 } 
	return x * factorial(x-1) 
}

func lFactorial(x int) int {
	total := 1 
	for ; 0 < x ; x-- {
		total *= x 
    }
	return total 
}
```

# 재귀와 스택프레임에 대한 참고자료   

https://www.youtube.com/watch?v=BNeOE1qMyRA