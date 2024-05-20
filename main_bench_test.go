package main

import (
	"math/rand/v2"
	"testing"
)

var maxN = 1000
var K = 1000

func createList(n int) *Node {
	root := &Node{rand.Int(), nil}
	
	prev := root
	for i := 0; i < n; i++ {
		cur := &Node{rand.Int(), nil}
		prev.Next = cur
		prev = cur
	}
	
	return root
}

func createLists(n, k int) []*Node {
	res := make([]*Node, 0, k)
	
	for i := 0; i < k; i++ {
		res = append(res, createList(n))
	}
	
	return res
}

func BenchmarkMergeKListsHeapWithCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		lists := createLists(maxN, K)
		b.StartTimer()
		mergeKListsHeapWithCopy(lists)
	}
}

func BenchmarkMergeKListsHeapInPlace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		lists := createLists(maxN, K)
		b.StartTimer()
		mergeKListsHeapInPlace(lists)
	}
}

func BenchmarkMergeKListsWithCopy(b *testing.B) {
	
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		lists := createLists(maxN, K)
		b.StartTimer()
		mergeKListsWithCopy(lists)
	}
}

func BenchmarkMergeKListsInPlace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		lists := createLists(maxN, K)
		b.StartTimer()
		mergeKListsInPlace(lists)
	}
}
