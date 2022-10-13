package main

import (
	"sync"
)

func criticalSection(m2 *sync.Mutex) {
	m2.Unlock()
}
func main() {
	var m1 sync.Mutex
	m1.Lock()
	criticalSection(&m1)
	m1.Lock()
	criticalSection(&m1)
}
