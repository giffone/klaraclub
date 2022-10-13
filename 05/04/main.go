package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var (
		a = int64(2)
		r = atomic.AddInt64(&a, 1)
	)
	fmt.Println(a, r)
}
