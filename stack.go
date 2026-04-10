package main

import "fmt"

const size int = 3

// LIFO
type stack struct {
	ptr   int
	array [size]int
}

func (s *stack) isEmpty() bool {
	return s.ptr == -1
}

func (s *stack) isFull() bool {
	return s.ptr == (size - 1)
}

func (s *stack) Push(value int) {
	if s.isFull() {
		return
	}
	s.ptr++
	s.array[s.ptr] = value
}

func (s *stack) Pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	value := s.array[s.ptr]
	s.ptr--
	return value, true
}

func (s *stack) Peek() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	return s.array[s.ptr], true
}

func main() {
	s := stack{
		ptr:   -1,
		array: [size]int{},
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.array)
	if v, ok := s.Pop(); ok {
		_ = v
	}
	if v, ok := s.Pop(); ok {
		_ = v
	}
	if v, ok := s.Pop(); ok {
		_ = v
	}
	s.Push(4)
	s.Push(4)
	fmt.Println(s.array)
}
