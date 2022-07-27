# 문자열

문자의 집합을 문자열이라고 한다.   
Go에서는 string이라 표기하고, 바이트 시퀀스이다.  
따라서 바이트 슬라이스로 변환할 수 있다. (아래는 변환 샘플)  

```go
func StringToBytes() {
	sample := "Hello, Hank"
	fmt.Printf("%T\n", sample) // string
	
	bSample := []byte(sample) 
	fmt.Println(bSample) // [72 101 108 ... ]
	fmt.Printf("%T\n", bSample) // []uint8
}
```

여기서 byte는 왜 uint8 일까?  
golang의 공식 스펙에 따르면 byte는 uint8의 alias기 떄문에  
byte와 uint8은 같은 자료형이라 볼 수 있다.  
