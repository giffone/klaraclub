package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a = int64(2)
	atomic.CompareAndSwapInt64(&a, 2, 1)
	fmt.Println(a)
}
