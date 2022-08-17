# Go Lint

Go Lint 는 어색한 코딩 스타일을 보고해준다.    
go v1.16은 `go get -u github.com/golang/lint/golint`  
go v1.18은 `go install github.com/golang/lint/golint@latest`  
로 설치한다.

사용 명령어는 `golint <filename>` 이다.

참고로,
* gofmt
    * 고 코드를 형식화 해준다.
* go vet
    * 의심스러운 구조체를 보고한다. 