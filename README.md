# googleAuth
谷歌身份验证器

# usage

```go
package main

import (
	"fmt"
)

func main() {

	auth := NewGoogleAuth()
	secret := auth.GetSecret()

	code, err := auth.GetCode(secret)
	url := auth.GetQrcodeUrl("gotest", secret)

	fmt.Println(secret, code, url, err)

	var input string

	for {
		fmt.Scanf("%s", &input)
		fmt.Println(input)

		b, err := auth.VerifyCode(secret, input)
		if err != nil || b == false{
			fmt.Println("failed")
		}else {
			fmt.Println("successful")
		}
	}

}
```

​	