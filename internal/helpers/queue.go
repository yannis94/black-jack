package helpers

import (
	"errors"
)

type node struct {
    Data interface{}
    prev *node
}

type Queue struct {
    Length int
    head *node
    tail *node
}

func NewQueue() *Queue {
    return &Queue{ Length: 0 }
}

func (q *Queue) Enqueue(data interface{}) {
    newNode := &node{ Data: data }


    if q.Length == 0 {
        q.head = newNode
        q.tail = newNode
        q.Length++

        return
    }

    q.tail.prev = newNode
    q.tail = newNode

    q.Length++

    return 
}

func (q *Queue) Dequeue() (interface{}, error) {
    if q.Length == 0 {
        return nil, errors.New("The queue is empty")
    }

    q.Length--

    n := q.head
    q.head = n.prev
    n.prev = nil

    return n.Data, nil
}

func (q *Queue) Peek() (interface{}, error) {
    if q.Length == 0 {
        return nil, errors.New("The queue is empty")
    }

    return q.head.Data, nil
}
