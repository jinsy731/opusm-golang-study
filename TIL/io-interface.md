# IO Interface

IO Interface 에는 Reader와 Writer가 있다.   
이 녀석들은 기본 입출력 인터페이스(`io.Writer`, `io.Reader`) 인데,  
Go 언어로 프로그래밍하면 자주 접하게 되는 녀석들이니 꼭 알아두어야 한다고 생각한다.  

Go 언어는 바이트(bytes) 를 사용하여 작업하기 위해 만들어진 프로그래밍이란 말도 있다.  
그래서 Writer, Reader 인터페이스는 바이트로 작업할 때 사용되는 기본적인 연산이다.  

이 인터페이스들은 웹 커넥션, 파일, 인메모리 슬라이스 등등 모든 표준 라이브러리에 걸쳐 구현된다.

## Spec

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}

type Writer interface {
        Write(p []byte) (n int, err error)
}
```

> 어떤 구조체든 매개변수로 바이트 슬라이스를 받고, 
> 정수와 에러 값을 리턴하는 Read 함수를 가지고 있으면 io.Reader 인터페이스를 따른다고 할 수 있습니다. 
> 마찬가지로 매개변수로 바이트 슬라이스를 받고 정수와 에러 값을 리턴하는 Write 함수를 가지고 있으면 
> io.Writer 인터페이스를 따른다고 할 수 있습니다.  
> 출처: http://pyrasis.com/book/GoForTheReallyImpatient/Unit50



# Encoder / Decoder

## Spec 

* type Decoder
  * func NewDecoder(r io.Reader) *Decoder
  * func (dec *Decoder) Buffered() io.Reader
  * func (dec *Decoder) Decode(v interface{}) error
  * func (dec *Decoder) More() bool
  * func (dec *Decoder) Token() (Token, error) 
  * func (dec *Decoder) UseNumber() 
* type Encoder
  * func NewEncoder(w io.Writer) *Encoder 
  * func (enc *Encoder) Encode(v interface{}) error
  * func (enc *Encoder) SetEscapeHTML(on bool) 
  * func (enc *Encoder) SetIndent(prefix, indent string) 

위 녀석들도 역시 Writer, Reader 인터페이스를 기반으로 구현되는데,  
입력이 웹 커넥션인지, 파일인지에 관계 없이 데이터를 클라이언트나 프로그램으로 돌려보낸다.  



## Encoder

어떤 데이터를 인코딩하는 경우에 `NewEncoder` 를 얻고, 그 인코더에 `Writer` 인터페이스를 제공한다.  
여기서 데이터는 __스트림__이고, `Encoder` 는 스트림을 __Json__ 문자열로 바꾼다.   
`Marshal` 과 비슷하다!

## Decoder 

`디코더(Decoder)` 는 입력에서 Json 으로 인코딩 된 값을 읽고 그것을 `v` 가 가리키는 값에 저장한다.  
그럼 디코더는 `Unmarshal` 과 비슷하다.  


# References

* https://godoc.org/encoding/json
* https://jeonghwan-kim.github.io/dev/2019/01/18/go-encoding-json.html
* http://pyrasis.com/book/GoForTheReallyImpatient/Unit50