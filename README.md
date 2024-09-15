# A Go client to [https://www.timeapi.io/](https://www.timeapi.io/)

```go
package main

import (
	"fmt"
	"time"

	"github.com/xaionaro-go/timeapiio"
)

func main() {
	t, err := timeapiio.Now()
	if err != nil {
		panic(err)
	}

	fmt.Printf("the difference between local and remote clock is %v\n", time.Since(t))
}
```

```
xaionaro@void:~/go/src/github.com/xaionaro-go/timeapiio$ go run ./example/main.go
the difference between local and remote clock is 14.386375ms
```