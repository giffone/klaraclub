package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	initSomething := func() {
		fmt.Println("Something has been init")
	}
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(initSomething)
			done <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
