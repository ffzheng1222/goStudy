package main

import (
	"fmt"
	"sync"
)

type SliceNum []int

func NewSlice() SliceNum {
	return make(SliceNum, 0)
}

func (s *SliceNum) Add(elem int) *SliceNum {
	fmt.Printf("Add start: SliceNum = %v\n", *s)
	*s = append(*s, elem)
	fmt.Printf("Add end: SliceNum = %v\n", *s)
	return s
}

func main() {
	once := sync.Once{}
	s := NewSlice()

	once.Do(func() {
		s.Add(16)
		//s.Add(18)
		//s.Add(20)
	})

	once.Do(func() {
		s.Add(18)
	})

	once.Do(func() {
		s.Add(20)
	})
}
