package main

import (
	"sync"
	"testing"
)

var (
	// Создаем указатель на слайс байт
	obj []byte
	// Создаем пул для слайсов байт
	bytePool = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1024)
			return &b
		},
	}
)

func BenchmarkMake(b *testing.B) {
	for N := 0; N < b.N; N++ {
		obj = make([]byte, 1024)
		_ = obj
	}
}
func BenchmarkBytePool(b *testing.B) {
	for N := 0; N < b.N; N++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
}
