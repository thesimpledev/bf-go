// Package stack holds our stack
package stack

import "sync"

type Stack struct {
	items []int
	mu    sync.Mutex
}

func New() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

func (s *Stack) Push(v int) {
	s.mu.Lock()
	s.items = append(s.items, v)
	s.mu.Unlock()
}

func (s *Stack) Pop() (int, bool) {
	s.mu.Lock()
	if len(s.items) == 0 {
		return 0, false
	}
	pos := len(s.items) - 1
	v := s.items[pos]
	s.items = s.items[:pos]
	s.mu.Unlock()
	return v, true
}
