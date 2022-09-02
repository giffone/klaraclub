package main

import (
	"fmt"
	"sync"
)

func multiply() {
	mm := [][]int{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
	}

	wg1, wg2 := sync.WaitGroup{}, sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for i := 0; i < len(mm); i++ {
			wg1.Add(1)
			go func(i int) {
				defer wg1.Done()
				for j := 0; j < len(mm[i]); j++ {
					mm[i][j] *= 2
				}
			}(i)
		}
	}()

	wg2.Wait()
	wg1.Wait()
	// time.Sleep(time.Second)

	for _, m := range mm {
		fmt.Print("{")
		for _, v := range m {
			fmt.Printf("%d,", v)
		}
		fmt.Println("}")
	}

}
