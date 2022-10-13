package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	first()
	second()

}

func first() {
	var (
		a = int64(2)
		r = atomic.SwapInt64(&a, 1)
	)
	fmt.Println(a, r)
}

func second() {
	var (
		a = int64(2)
	)
	atomic.SwapInt64(&a, 1)
	fmt.Println(a)
}
