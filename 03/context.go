package main

import (
	"context"
	"fmt"
)

func a2() {
	fmt.Println("a")
	cancel()
	wg.Done()
}

func b2() {
	<-ctx.Done()
	fmt.Println("b")
	cancel2()
	wg.Done()
}

func c2() {
	<-ctx2.Done()
	fmt.Println("c")
	wg.Done()
}

var ctx, cancel = context.WithCancel(context.Background())
var ctx2, cancel2 = context.WithCancel(context.Background())

func withContext() {
	fmt.Println("with context")
	wg.Add(3)
	go a2()
	go b2()
	go c2()
	wg.Wait()
}
