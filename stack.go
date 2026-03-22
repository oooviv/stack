package main

import "fmt"

const size int = 4

type stack struct {
	left  int
	right int
	array [size]int
}

func (s *stack) isEmpty() bool {
	return s.left == -1
}

func (s *stack) isFull() bool {
	return s.right == size-1
}

func (s *stack) Push(v int) bool {
	if s.isFull() {
		return false
	}
	if s.left == -1 {
		s.left = 0
	}
	s.right++
	s.array[s.right] = v
	return true
}

func (s *stack) Pop() (int, bool) {
	value := 0
	if s.isEmpty() {
		return 0, false
	}
	if s.left == s.right {
		value = s.array[s.right]
		s.array[s.right] = 0
		s.left = -1
		s.right = -1
	} else {
		value = s.array[s.right]
		s.array[s.right] = 0
		s.right--
	}
	return value, true
}

func (s *stack) Peek() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	return s.array[s.right], true
}

func main() {
	s := stack{
		left:  -1,
		right: -1,
	}
	_ = s

	if s.Push(1) {
		fmt.Println("Added:", 1)
	}
	if s.Push(2) {
		fmt.Println("Added:", 2)
	}
	if s.Push(3) {
		fmt.Println("Added:", 3)
	}
	if s.Push(4) {
		fmt.Println("Added:", 4)
	}
	if s.Push(5) {
		fmt.Println("Added:", 5)
	}
	fmt.Println(s.array)

	if v, ok := s.Pop(); ok {
		fmt.Println("Deleted:", v)
	}
	if v, ok := s.Pop(); ok {
		fmt.Println("Deleted:", v)
	}
	if v, ok := s.Pop(); ok {
		fmt.Println("Deleted:", v)
	}
	if s.Push(5) {
		fmt.Println("Added:", 5)
	}
	fmt.Println(s.array)
}
