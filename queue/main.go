package main

import (
	"errors"
	"fmt"
)

var show = true

type Queue struct {
	nodes []int
}

func createQueue() *Queue {
	return &Queue{}
}

func (q *Queue) push(nodeVal int) error {
	if show {
		fmt.Printf("Pushed %d to Queue\n", nodeVal)
	}
	q.nodes = append(q.nodes, nodeVal)
	return nil
}

func (q *Queue) dequeue() (int, error) {
	if len(q.nodes) != 0 {
		res := q.nodes[0]
		q.nodes = q.nodes[1:]
		if show {
			fmt.Printf("Dequeued %d from Queue\n", res)
		}
		return res, nil
	} else {
		return 0, errors.New("The Queue is Empty")
	}
}

func (q *Queue) clearQ() error {
	for range q.nodes {
		q.dequeue()
	}
	return nil
}

func main() {
	q := createQueue()
	q.push(12)
	q.push(31)
	q.push(222)
	q.clearQ()
	res, err := q.dequeue()
	if res == 0 {
		fmt.Println(err)
	}
}
