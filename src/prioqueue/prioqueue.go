package prioqueue

import (
	"chizu-ru/node"
	"errors"
)

type queueElement struct {
	info     *node.Node
	priority float64
}

type PrioQueue struct {
	elements []*queueElement
}

func New() *PrioQueue {
	pq := new(PrioQueue)
	pq.elements = make([]*queueElement, 0)

	return pq
}

func (pq *PrioQueue) Enqueue(n *node.Node, prio float64) {
	if pq.IsEmpty() {
		pq.elements = append(pq.elements, &queueElement{n, prio})
		return
	}

	for i, el := range pq.elements {
		if el.priority > prio {
			temp := pq.elements[:i]
			temp = append(temp, &queueElement{n, prio})
			pq.elements = append(temp, pq.elements[i:]...)
			return
		}
	}

	pq.elements = append(pq.elements, &queueElement{n, prio})
}

func (pq *PrioQueue) Dequeue() (*node.Node, error) {
	if pq.IsEmpty() {
		return nil, errors.New("queue kosong")
	}

	temp := pq.elements
	pq.elements = pq.elements[1:]
	return temp[0].info, nil
}

func (pq *PrioQueue) IsEmpty() bool {
	return len(pq.elements) == 0
}
