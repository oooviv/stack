package main

import "fmt"

const size int = 4

type stack struct {
	front int
	array [size]int
	reare int
}

// IsEmpty возвращает true если стек пустой
func (s *stack) IsEmpty() bool {
	return s.front == -1
}

// IsFull возвращает true если стек полный
func (s *stack) IsFull() bool {
	return s.reare == size-1
}

func (s *stack) Push(v int) bool {
	if s.IsFull() {
		return false
	}
	if s.front == -1 {
		s.front = 0
		s.reare++
		s.array[s.reare] = v
	} else {
		s.reare++
		s.array[s.reare] = v
	}
	return true
}

func (s *stack) Pop() (int, bool) {
	value := 0
	if s.IsEmpty() {
		return 0, false
	}
	if s.reare == 0 {
		value = s.array[s.reare]
		s.front = -1
		s.reare = -1
	} else {
		value = s.array[s.reare]
		s.reare--
	}
	return value, true
}

func (s *stack) Peek() int {
	value := 0
	if s.IsEmpty() {
		return value
	}
	value = s.array[s.reare]
	return value
}

func main() {
	s := stack{
		front: -1,
		array: [size]int{},
		reare: -1,
	}
	fmt.Println()
	fmt.Println(s.Peek())
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
	fmt.Println(s.Peek())
	fmt.Println()
	if j, ok := s.Pop(); ok {
		fmt.Println("Deleted:", j)
	}
	if j, ok := s.Pop(); ok {
		fmt.Println("Deleted:", j)
	}
	fmt.Println(s.Peek())
	fmt.Println()
	if s.Push(10) {
		fmt.Println("Added:", 10)
	}
	fmt.Println(s.Peek())
	if s.Push(20) {
		fmt.Println("Added:", 20)
	}
	fmt.Println(s.Peek())
	fmt.Println()
}
