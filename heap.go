package main

import (
	"container/heap"
)

type Node struct {
	Value int
	Next *Node
}


type Heap []*Node

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(*Node))
}

func (h *Heap) Pop() any {
	old := *h
	l := len(old)
	value := old[l-1]
	*h = old[0:l-1]
	return value
}

///////////////////////////////////////////////////////////////
// InPlace solution
///////////////////////////////////////////////////////////////

func mergeKListsHeapInPlace(lists []*Node) *Node {
	
	if len(lists) == 0 {
		return nil
	}
	
	heapLists := Heap(lists)
	heap.Init(&heapLists)
	
	var root *Node
	var prev *Node
	
	for heapLists.Len() > 0 {
		// pop min
		cur := heap.Pop(&heapLists).(*Node)
		
		// push next
		if cur.Next != nil {
			heap.Push(&heapLists, cur.Next)
		}
		
		if root == nil {
			root = cur
		}
		
		if prev != nil {
			prev.Next = cur
		}
		
		prev = cur
	}
	
	return root
}

///////////////////////////////////////////////////////////////
// Copy solution
///////////////////////////////////////////////////////////////
func mergeKListsHeapWithCopy(lists []*Node) *Node {
	
	if len(lists) == 0 {
		return nil
	}
	
	heapLists := make(Heap, 0, len(lists))
	heapLists = append(heapLists, lists...)
	
	heap.Init(&heapLists)
	
	var root *Node
	var prev *Node
	
	for heapLists.Len() > 0 {
		// pop min
		minNode := heap.Pop(&heapLists).(*Node)
		
		// push next
		if minNode.Next != nil {
			heap.Push(&heapLists, minNode.Next)
		}
		
		cur := &Node{minNode.Value, nil}
		
		if root == nil {
			root = cur
		}
		
		if prev != nil {
			prev.Next = cur
		}
		
		prev = cur
	}
	
	return root
}
