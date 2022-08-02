# Interface

`인터페이스( Interface )` 는 Go에서 행동과 다형성을 구현하는 방법이다.  
인터페이스가 구현되는 순간, 특정 값은 하나 이상의 타입이 될 수 있다.  

```go
package main 

import "fmt"

type preson struct {
	last string
	first string
}

type secretAgent struct { 
	person 
	patner string 
	ltk bool 
}

type human interface {      // human 인터페이스는 타입이고, secretAgent 타입의 speak()를 포함하고 있다. 
	speak() 
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}

func main() {
	sa1 := secretAgent {
		person: person {
			first: "Kim",
			last: "HanKyul",
        },
		panter: "SeungHyun",
		ltk: true,
    }
	
	sa1.speak()
}
```

위 샘플에서 `human interface` 가 새로 정의되어 있다. Go 언어에서는 모든게 타입이고 인생도 타입이라고 했다. 따라서 `interface` 도 당연히 타입인데, 이 타입은 현재 `speak()` 라고하는 `secretAgent` 타입의 메소드를 갖고 있다.  
그리고 메인 함수에서는 식별자 `sa1` 이 `type secretAgent` 를 할당 받았다. 할당된 값은 `human interface` 의 `speak()` 메소드를 사용할 수 있으므로 `sa1` 에는 두 개 이상의 타입이 존재한다.

어디서 많이 본 구조 아닌가?  
OOP에서의 인터페이스도 행동과 다형성을 정의한다.  
정의된 인터페이스는 클래스로 상속해서 공통된 사항을 오버라이딩해 구현하기도 한다.  

Go 언어에서는 비록 `extends` 나 `implements` 같은 상속 키워드는 없지만, `interface` 와 `reciever` 를 사용하여 OOP를 구현할 수 있다.  

```go
package main

import "fmt"

type human interface { 
	speak() 
}

type person struct { 
	first string
	last string 
}

type secretAgent struct { 
	person 
	patner string 
	ltk bool
}

func (p person) speak() {
	fmt.Println(p.first, p.last) 
}

func (s secretAgent) speak() {
	fmt.Println(s. first, s.last)  
}

func helloAgent(h human) {
	fmt.Println("Hello" , h) 
}

func main() {
	sa1 := secretAgent { 
		person : person { 
			first: "Kim",
			last: "HanKyul", 
        },  
		ltk: true, 
		patner: "SeungHyun",
    }
	
	p1 := person {
		first: "James",
		last: "Bond", 
    }
	
	p1.speak()
	sa1.speak()
	helloAgent(p1)
	helloAgent(sa1)
}
```

위 샘플에서 보듯이, `interface` 는 `speak()` 라는 행동을 정의하고 있고, 이를 `person`, `secretAgent` 에서 사용할 수 있도록 리시버 매개변수를 통해서도 정의했다. 
그리고 `helloAgent()` 에서는 `human interface` 를 매개변수로 받는데, 메인 함수에서는 `p1` 과 `sa1` 두 타입 모두를 전달하고 있다.  

왜냐하면 위에서도 설명 했듯이, Go는 타입이 전부인 언어기 때문이다. 타입으로 모든걸 표현한다.  
human interface 는 speak() 를 정의한 타입이고, speak() 는 person 과 secretAgent 가 사용하므로 person, secretAgent 의 타입은 human interface 라 할 수 있다. 
따라서 human interface 타입은 person, secertAgent 타입을 취할 수 있다.  

이것이 Go에서의 polymorphism(다형성) 이다. 다형성은 의미 그대로 많이 변한다는 뜻이다.  
human 이 person 이 되고, secretAgent 가 될 수 있는 것처럼 말이다. 