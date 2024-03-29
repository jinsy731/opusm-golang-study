# Go는 OOP인가?  

맞기도 하고 아니기도 하다.  
Go에는 타입, 메소드가 있고 객체지향 스타일을 허용한다.  
그러나 타입 계층이 없다.  

Go의 인터페이스 개념은 우리가 생각하기에 더 일반적으로 사용하기 쉬운 다른 접근 방식을 제공한다.  

서브클래싱과 유사하지만, 동일하지 않은 것을 제공하기 위해 다른 타입에 타입을 포함하는 방법도  
있다.  

또한 Go의 메소드는 C++ 또는 Java보다 더 일반적이다.  
모든 종류의 데이터에 대해 정의할 수 있으며, 일반적인 그리고 방식되지 않은 정수와 같은  
기본 제공 타입을 포함한다. 그것들은 구조체 클래스로 제한된 구조체가 아니다.

또한 타입 계층이 없기 때문에 Go의 객체는 C++ 또는 Java와 같은 언어보다 훨씬 가볍다.  

그렇다면 Go의 OOP스러운 점은 뭘까? 

* 캡슐화 되어있다. 
* 상태, 필드를 가질 수 있다.
* 동작(behavior)을 가질 수 있다. 동작에 메소드를 넣을 수 있고 뺄 수도 있다.
* 재사용이 가능하다. 
* 임베디드 타입에 대한 상속이 있고 다른 타입에 타입을 임베드 할 수 있다.  
* 인터페이스로 다형성이 가능하고, 오버라이딩 가능하다. 

전통적인 OOP는 뭘까? 

* 객체의 타입을 설명하는 데이터 구조인 클래스가 있다.  
* 클래스로 인스턴스, 객체를 만들 수 있다.  
* 클래스는 상태와 필드를 모두 소유한다. 
* 동작을 나타내는 메소드가 있다. 

Go와 전통적인 OOP의 차이는 뭘까?  

* Go에서는 private, public이라는 개념이 없고, exported, unexported 또는 viewable, not viewable 쓴다.  
* Go에서는 클래스가 아닌 타입을 생성한다.  
* 인스턴스화 하지 않고 타입의 값을 생성한다.  