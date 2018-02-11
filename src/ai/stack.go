package ai

import "fmt"

type intNode struct {
	data int
	next *intNode
}

type Stack struct {
	head *intNode
}

func NewStack() *Stack {
	s := &Stack{nil}
	return s
}

func (s *Stack) Push(data int) {
	n := &intNode{data: data, next: s.head}
	s.head = n
}

func (s *Stack) Pop() (int, bool) {
	n := s.head
	if s.head == nil {
		return 0, false
	}
	s.head = s.head.next
	return n.data, true
}

func (s *Stack) String() string {
	st := ""
	now := s.head
	for now != nil {
		st += fmt.Sprint(now.data)
		now = now.next
	}
	return st
}
