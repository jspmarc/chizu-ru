// Package prioqueue lirary priority queue dengan elementnya adalah pointer
// sudut pada graf. Bisa ada beberapa elemen yang sama dengan priority berbeda.
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
	a := new(queueElement)
	a.info = n
	a.priority = prio

	if pq.IsEmpty() {
		pq.elements = append(pq.elements, a)
		return
	}

	for i, el := range pq.elements {
		if el.priority > prio {
			temp := make([]*queueElement, 0)
			temp = append(temp, pq.elements[:i]...)
			temp = append(temp, a)
			pq.elements = append(temp, pq.elements[i:]...)
			return
		}
	}

	pq.elements = append(pq.elements, a)
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

// ContainsLEQ memeriksa apakah di queue sudah ada elemen yang sama dengan
// priority lebih rendah atau sama
func (pq *PrioQueue) ContainsLEQ(n *node.Node, prio float64) bool {
	for _, e := range pq.elements {
		if e.info == n && e.priority <= prio {
			return true
		}
	}

	return false
}

// RemoveEl menghapus elemen n dari queue
func (pq *PrioQueue) RemoveEl(n *node.Node) {
	tmp := make([]*queueElement, 0)
	for i, e := range pq.elements {
		if e.info == n {
			tmp = append(tmp, pq.elements[:i]...)
			tmp = append(tmp, pq.elements[i+1:]...)
			pq.elements = tmp
		}
	}
}

// Clear mengosongkan queue
func (pq *PrioQueue) Clear() {
	for !pq.IsEmpty() {
		pq.Dequeue()
	}
}
