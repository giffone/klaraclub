package main

import "fmt"

func reverseString(s []byte) {
	lS := len(s)

	for i := 0; i < lS/2; i++ {
		s[i], s[lS-1-i] = s[lS-1-i], s[i]
	}

	fmt.Println(string(s))
}
