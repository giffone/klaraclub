package main

func reverseString(s []byte) {
	lS := len(s)
	if lS < 2 {
		return
	}
	s[0], s[lS-1] = s[lS-1], s[0]
	reverseString(s[1 : lS-1])
}
