# Go workspace 

Golang은 팀의 효율성과 생산성을 위해 코드 형식까지 정해 놓았다.  
때문에 어떤 컴퓨터에서 작업하던지 정해진 약속을 따르게 함으로써 모든 게 어떻게 작동할지 예측할 수 있다.  

Workspace도 그런 규약 중에 하나이다.   
```
one directory  
|-bin - complie된 내용(실행파일 및 바이너리) 저장 
|-pkg - 압축된 파일(패키지 오브젝트) 저장 (컴파일링 절약을 위해 패키지가 변경되지 않는다면 아키이브를 이곳에 저장)
|-src - 소스코드, 프로젝트 폴더 또는 리포지토리에 관한 네임스페이싱, 패키지 관리  
|--github.com  
|---<github.com username>  
|----directory with code for project / repo  
|----directory with code for project / repo  
|----directory with code for project / repo  
```
workspace는 위와 같은 구조로 되어 있다. 

# GOPATH 설정 방법

콘솔에서 go env 명령어를 입력하는 아래와 같은 모습을 볼 수 있는데,

```
GO111MODULE=""
GOARCH="arm64"
GOBIN=""
GOCACHE="/Users/kimh4nkyul/Library/Caches/go-build"
GOENV="/Users/kimh4nkyul/Library/Application Support/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="arm64"
GOHOSTOS="darwin"
...
```
여기서 사용자가 만든 workspace를 사용하고 싶다면 GOPATH를 변경해야 한다.  
GOPATH는 workspace 경로가 설정된 환경 변수로 사용자 계정 디렉토리 밑에 go가 기본 설정되어 있다.  
(예를 들어, /Users/kimh4nkyul/go)  

필자는 Mac M1에서 Oh-My-Zsh을 사용하고 있기 때문에 아래와 같은 방법으로 GOPATH를 변경했다.  
(Linux나 Windows를 사용하고 있다면 구글신께 물어보길 바란다. 😄)
  
```
cd $HOME    // 현재 사용자 계정의 홈 디렉토리로 이동
vi .zshrc   // 현재 사용자 환경변수 설정(전역 사용자는 /etc/zshrc에서 설정)
export GOPATH="$HOME/<workspace로 지정할 디렉토리>" // 설정후 저장
source .zshrc
```

위 과정을 통해 GOPATH에 대한 환경 변수를 설정했다면, go env 명령어로 설정이 잘 되었는지 확인한다.  
```
go env | grep GOPATH
GOPATH="/Users/kimh4nkyul/IdeaProjects/opusm-golang-study/TIL/go-workspace"
```