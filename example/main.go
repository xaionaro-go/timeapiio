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
