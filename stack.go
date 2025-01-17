package minequery

import (
	"errors"
)

var errStackEmpty = errors.New("stack is empty")

type stack []interface{}

func (s *stack) Push(value interface{}) { *s = append(*s, value) }

func (s *stack) Pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errStackEmpty
	}
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret, nil
}
