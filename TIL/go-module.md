# Go Module

Go Module은 Go로 작성한 프로그램에서 사용할 수 있는 패키지 모음을 의미한다.  
1.11 버전에서 소개되고, 1.13 버전에 기능이 완성되었으며 1.16 버전에서 기본 사양이 되었다.  
Go Module은 아래와 같은 특징을 갖고 있다.  
* 한 개의 모듈은 다수의 패키지를 포함할 수 있다.  
* 모듈을 통해 패키지 종속성을 관리할 수 있다.  
* 모듈은 패키지 관리 시스템으로써 활용된다.  
* 모듈은 패키지를 트리 형식으로 관리한다.  
* 모듈의 이름은 유니크해야 한다.  

## Go Module 생성 방법
```
Usage : go mod init <URL for namespacing> 
go mod init example.com/test
라던가...
go mod init github.com/KimH4nkyul/go-test-module

<Note>
Name Spacing(네임 스페이싱) : 
    Go에서의 네임 스페이싱은 Java의 Project명과 같이 '도메인명'을 의미한다.
    따라서, 프로젝트나 조직을 나타내는 특수한 이름을 사용한다. 
    보통은 Git Repository의 명을 사용한다.
    (예)
```

위의 생성 방법에서도 알 수 있듯이, Go Module(go mod)은  
새로운 모듈을 만들 수 있고, `go.mod` 파일을 정의, 초기화할 수 있다. (`go init`) 

## Go Module 관련 명령어

Go Module과 관련된 명령어가 있는데  
`go build`, `go test`, `go list -m all`, `go get` 다.   

* `go build`
* `go test` 
  * go build와 go test는 패키지 구축 명령어로 사용된다. 
  * go.mod에 새로운 종속성을 추가하는데 사용된다. 
* `go list -m all` : 
  * 현재 모듈의 종속성을 프린트한다. 
* `go get` : 
  * 새로운 종속성을 추가하거나 필요한 종속성의 버전을 변경한다. 

## Go Module에서 사용하지 않는 종속성 제거

```
go mod tidy 
```

## Go Module로 사용할 Repository 다운로드 방법

```
go get -d <URL>
go get -d github.com/KimH4nKyul/go-test-module/module1
go get -d github.com/KimH4nKyul/go-test-module/...

<Note>
go get -d 사용할 때, 맨 뒤에 ...을 붙이면  
관련된 모든 패키지를 재귀적으로 받겠다는 의미이다. 
```

현재 go get은 <b>1.17버전</b> 부터 Deprecated 되었다.  
따라서 go install을 통해 패키지를 다운로드 받으면 된다.
go install로 패키지를 다운받을 경우에는 반드시 `@latest` 같이 버전을 명시해야 한다.

go install은 패키지를 다운받고 컴파일하는 과정이 포함되어 있다. 

```
go install github.com/KimH4nKyul/go-test-module/module1@latest
```

설치가 완료된 패키지는 `$GOPATH/pkg` 에 위치할 것이다. 