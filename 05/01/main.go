package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const count = 1000

func main() {
	first()
	fmt.Println("---------")
	second()
	fmt.Println("---------")
	third()
}

func first() {
	var (
		counter int32
		ch      = make(chan struct{}, count)
	)

	for i := 0; i < count; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			// counter += 1
			ch <- struct{}{}
		}()
	}

	time.Sleep(2 * time.Second)
	close(ch)

	i := 0

	for range ch {
		i += 1
	}

	fmt.Println(counter)
	fmt.Println(i)
}

func second() {
	var (
		counter int32
		ch      = make(chan struct{}, count)
	)

	mut := sync.Mutex{}

	for i := 0; i < count; i++ {
		go func() {
			mut.Lock()
			counter++
			mut.Unlock()
			ch <- struct{}{}
		}()
	}

	time.Sleep(2 * time.Second)
	close(ch)

	i := 0

	for range ch {
		i += 1
	}

	fmt.Println(counter)
	fmt.Println(i)
}

func third() {
	wg := sync.WaitGroup{}
	var (
		counter int32
		ch      = make(chan struct{}, count)
	)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
			ch <- struct{}{}
		}()
	}
	wg.Wait()
	close(ch)

	i := 0

	for range ch {
		i += 1
	}

	fmt.Println(counter)
	fmt.Println(i)
}
