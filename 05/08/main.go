package main

import (
	"fmt"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

const count = 1000

func main() {
	var (
		counter int32
		eg      = errgroup.Group{}
	)

	for i := 0; i < count; i++ {
		i := i
		eg.Go(func() error {
			atomic.AddInt32(&counter, 1)
			if i%10 == 0 {
				return fmt.Errorf("error on calculate %d", i)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("counter:", counter)
}
