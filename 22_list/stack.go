package main

import (
	"container/list"
	"fmt"
)

// Stack 자료구조 정의
type Stack struct {
	v *list.List
}

// Push
func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

// Pop
func (s *Stack) Pop() interface{} {
	top := s.v.Back()
	if top != nil {
		return s.v.Remove(top)
	}
	return nil
}

// Stack 생성
func NewStack() *Stack {
	return &Stack{list.New()}
}

func main() {

	s := NewStack()

	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	fmt.Println()
	top := s.Pop()
	for top != nil {
		fmt.Print(top, "-> ")
		top = s.Pop()
	}

}
