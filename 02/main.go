package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("test2")
	if err != nil {
		log.Fatalln(err)
	}
	in := bufio.NewReader(file)
	// in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	option := strings.Split(line[:len(line)-1], " ")
	if len(option) != 2 {
		log.Fatalln("option not correct")
	}
	loop, err := strconv.Atoi(option[0])
	if err != nil {
		log.Fatalf("read loop-option: %s\n", err.Error())
	}
	find, err := strconv.Atoi(option[1])
	if err != nil {
		log.Fatalf("read find-option: %s\n", err.Error())
	}

	matrix := make([][]int, loop)
	y := 0
	for y < loop {
		line, err := in.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				line += "\n"
			} else {
				log.Fatalf("read matrix line: %s", err.Error())
			}
		}

		slice := strings.Split(line[:len(line)-1], " ")
		lSlice := len(slice)
		if lSlice != loop {
			log.Fatalf("matrix sides wrong, expect %d but got %d", loop, lSlice)
		}
		matrix[y] = make([]int, loop)
		for i := 0; i < lSlice; i++ {
			n := 0
			if slice[i] == "1" {
				n = 1
			}
			matrix[y][i] = n
		}
		y++
	}
	a := answer{}
	found := []int{find - 1}
	loop = 0
	for {
		for _, x := range found {
			buf := []int{}
			for i, n := range matrix[x] {
				if n != 0 {
					if !a.exist(i + 1) {
						a.vertex = append(a.vertex, i+1)
					}
					if i < x {
						continue
					}
					buf = append(buf, i)
				}
			}
			found = buf
		}
		if len(found) == 0 {
			break
		}
	}
	out := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(out, len(a.vertex))
	out.Flush()
}

type answer struct {
	vertex []int
}

func (a answer) exist(n int) bool {
	for _, v := range a.vertex {
		if v == n {
			return true
		}
	}
	return false
}
