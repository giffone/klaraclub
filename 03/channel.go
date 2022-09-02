package main

import (
	"fmt"
	"sync"
)

func a() {
	fmt.Println("a")
	ch1 <- struct{}{}
	wg.Done()
}

func b() {
	<-ch1
	fmt.Println("b")
	ch2 <- struct{}{}
	wg.Done()
}

func c() {
	<-ch2
	fmt.Println("c")
	wg.Done()
}

var ch1 = make(chan struct{})
var ch2 = make(chan struct{})
var wg sync.WaitGroup

func withChannel() {
	fmt.Println("with channel")
	wg.Add(3)
	go a()
	go b()
	go c()
	wg.Wait()
}
