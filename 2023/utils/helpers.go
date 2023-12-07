package utils

import "errors"

type Stack struct {
	top      int
	Contents []string
}

func (s *Stack) Push(elem string) {
	if s.Contents == nil {
		s.Contents = []string{}
	}
	s.top++
	s.Contents = append(s.Contents, elem)
}

func (s *Stack) Pop() (string, error) {
	if len(s.Contents) == 0 {
		return "", errors.New("Stack empty")
	}
	last := s.top - 1
	elem := s.Contents[last]
	s.Contents = s.Contents[:last]
	s.top--
	return elem, nil
}

func (s *Stack) Peek() (string, error) {
	if len(s.Contents) == 0 {
		return "", errors.New("Stack empty")
	}
	return s.Contents[s.top-1], nil
}
