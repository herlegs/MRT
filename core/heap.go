package core

import "container/heap"

type Node struct {
	Cost    int
	Station *Station
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type NodeHeap struct {
	pq *PriorityQueue
}

func (h *NodeHeap) Push(node *Node) {
	heap.Push(h.pq, node)
}

func (h *NodeHeap) Pop() *Node {
	return heap.Pop(h.pq).(*Node)
}

func (h *NodeHeap) Len() int {
	return h.pq.Len()
}

func NewNodeHeap() *NodeHeap {
	return &NodeHeap{
		pq: &PriorityQueue{},
	}
}
