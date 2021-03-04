package main

import (
	"errors"
	"fmt"
)

var show = true

type Stack struct {
	nodes []int
}

func createStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(nodeVal int) error {
	if show {
		fmt.Printf("Pushed %d to Stack\n", nodeVal)
	}
	s.nodes = append(s.nodes, nodeVal)
	return nil
}

func (s *Stack) pop() (int, error) {
	if len(s.nodes) != 0 {
		res := s.nodes[len(s.nodes)-1]
		s.nodes = s.nodes[:len(s.nodes)-1]
		if show {
			fmt.Printf("Poppped %d from Stack\n", res)
		}
		return res, nil
	} else {
		return 0, errors.New("The Stack is Empty")
	}
}

func (s *Stack) clearS() error {
	for range s.nodes {
		s.pop()
	}
	return nil
}

func main() {
	s := createStack()
	s.push(12)
	s.push(31)
	s.push(222)
	s.push(221)
	s.push(220)
	s.push(219)
	s.clearS()
	res, err := s.pop()
	if res == 0 {
		fmt.Println(err)
	}
}
