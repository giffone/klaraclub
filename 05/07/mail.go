package main

import "sync"

type Set struct {
	sync.Map
}

func NewSet() *Set {
	return &Set{
		Map: sync.Map{},
	}
}
func (s *Set) Add(i int) {
	s.Map.Store(i, struct{}{})
}
func (s *Set) Has(i int) bool {
	_, ok := s.Map.Load(i)
	return ok
}

func main() {

}
