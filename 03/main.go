package main

import (
	"time"
)

func main() {
	withChannel()
	time.Sleep(100 * time.Millisecond)
	withContext()
	time.Sleep(100 * time.Millisecond)
	some()
}
