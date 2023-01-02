package main

import (
	"container/list"
	"fmt"
)

// Queue 정의
type Queue struct {
	v *list.List
}

// 요소 Push
func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

// 요소 Pop
func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

// 새로운 Queue 생성
func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {

	q := NewQueue()

	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	v := q.Pop()
	for v != nil {
		fmt.Print(v, "-> ")
		v = q.Pop()
	}

}
