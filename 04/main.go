package main

import "fmt"

func main() {
	fmt.Println("multiply")
	multiply()
	fmt.Println("\nreverse")
	b := []byte{'f', 'a', 'i', 'z', 'u', 'l', 'l', 'a'}
	reverseString(b)
	fmt.Println(string(b))
	fmt.Println("\nswap nodes in pairs")
	makeSwap()
}
