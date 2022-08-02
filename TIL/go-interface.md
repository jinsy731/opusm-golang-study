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
	fmt.Println("Hello" , h.first) 
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

좀 더 쉬운 비유를 들어 본다면.  
`human` 이라는 사람의 이상형이 `speak()` 할 줄 아는 사람인 거고,  
그 조건에 포함되는 사람인 `sa1`, `p1` 에게 "당신 정말 제 타입이네요. " 라고 말하는 상황인거다.  
그런데 인생을 살면서 이상형이 언제나 한결같은 적이 없었을 뿐. 🤣

# 인터페이스 단언(Assert)  

```go
package main 
import "fmt" 

// ... 

func helloAgent(h human) {
	switch h.(type) {
	case person:
		fmt.Println(h.(person).first)
	case secretAgent: 
		fmt.Println(h.(secretAgent).first)
	default:
		fmt.Println(h.first)
    }
}

// ...
```

당연히 다형성을 구현하는 개념인 만큼, `conversion` 가능하다.  
그러나 위 샘플에서 사용된 개념은 엄연히 다르고, 타입에 대한 단언 처리를 하는 것이다.  
전환 이라는 것은 어떤 값을 다른 타입으로 변경하는 것이다.  
위 샘플에서 구현된 방법은 어떤 변수의 타입에 대한 처리를 분기하는 `assertion` 혹은 `asserting` 이라는 기법이다.

# 익명 함수  

익명 구조체에 이어 익명 함수가 등장했다.  
익명 함수는 별다른 개념 설명 없이 샘플만 봐도 이해가 가능하다.  
그냥 이름이 없는 함수다.  

```go
package 

import "fmt" 

func main() {
	func(lang string) {
		fmt.Println("My favorite programming language is ", lang)
    }("Go")
}
```

익명 함수의 구조는 `func(paramters) { code } (arguments)` 이며,  
익명 함수의 존재 이유는 재사용할 필요가 없고, 호출되고 종료되는 그 한 순간에서 메모리 관리에 대한 이점을 얻기 때문이다.  
(일반 함수는 언제든 메모리에서 여러번 또는 항상 존재할 경우가 있지만. )


