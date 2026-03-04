package main

import "fmt"

type stack struct {
	array []int
}

func (s *stack) IsEmpty() bool {
	return len(s.array) == 0
}

// func (s *stack) IsFull() bool {
// 	return len(s.array) == cap(s.array)
// }

func (s *stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.array[len(s.array)-1], true
}

func (s *stack) Push(value int) {
	// if !s.IsFull() {
	// 	s.array = append(s.array, value)
	// }
	s.array = append(s.array, value)
}

func (s *stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	value := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return value, true
}

func main() {

	st := stack{
		array: make([]int, 0, 10),
	}
	fmt.Println()
	st.Push(0)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)
	st.Push(6)
	st.Push(7)
	st.Push(8)
	st.Push(9)
	fmt.Println(st.array)
	if v, ok := st.Peek(); ok {
		fmt.Println("Peek -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	fmt.Println()
	st.Push(0)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)
	st.Push(6)
	st.Push(7)
	st.Push(8)
	st.Push(9)
	st.Push(10)
	st.Push(11)
	st.Push(12)
	fmt.Println(st.array)
	if v, ok := st.Peek(); ok {
		fmt.Println("Peek -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop -->", v)
	}
	fmt.Println()
}
