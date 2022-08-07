# bcrypt 

우리가 어딘가에 패스워드를 저장하려고 할 때,  
`bcrypt` 는 __암호화된__ 패스워드를 저장하는 탁월한 솔루션이 될 수 있다.  

`bcrypt` 는 적응형 해싱 알고리즘이다.  
(https://pkg.go.dev/golang.org/x/crypto/bcrypt)

```go
package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error: ", err) 
    }
	
	loginPwd := "password1234" 
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd)) 
	if err != nil {
		fmt.Println("You can't login") 
		return 
    }
	fmt.Println("You logged in")
}
```