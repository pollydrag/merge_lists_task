package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	Lists [][]int
	Expected []int
}{
	{
		[][]int{[]int{2,5}, []int{1,4,6}, []int{1,3}, []int{2,5,7}},
		[]int{1,1,2,2,3,4,5,5,6,7},
	},
	{
		nil,
		nil,
	},
	{
		[][]int{},
		nil,
	},
	{
		[][]int{[]int{1}, []int{1}, []int{1}},
		[]int{1,1,1},
	},
	{
		[][]int{[]int{1}, []int{1}, []int{}},
		[]int{1,1},
	},
}

func toNode(list []int) *Node {
	if len(list) == 0 {
		return nil
	}
	
	root := &Node{list[0], nil}
	
	prev := root
	for i := 1; i < len(list); i++ {
		cur := &Node{list[i], nil}
		prev.Next = cur
		
		prev = cur	
	}
	
	return root
}

func toArrayNodes(lists [][]int) []*Node {
	ret := make([]*Node, 0)
	
	for _, list := range lists {
		node := toNode(list)
		if node == nil {
			continue
		}
		ret = append(ret, toNode(list))
	}
	
	return ret
}

func toArray(n *Node) []int {
	if n == nil {
		return nil
	}
	
	array := make([]int, 0)
	for n != nil {
		array = append(array, n.Value)
		n = n.Next
	}
	
	return array
}

func TestMergeListsHeapInPlace(t *testing.T) {
	for _, test := range tests {
		lists := toArrayNodes(test.Lists)
		
		res := mergeKListsHeapInPlace(lists)
		assert.Equal(t, test.Expected, toArray(res))
	}
}

func TestMergeListsHeapWithCopy(t *testing.T) {
	for _, test := range tests {
		lists := toArrayNodes(test.Lists)
		listCopy := make([]*Node, len(lists))
		copy(listCopy, lists)
		
		res := mergeKListsHeapWithCopy(lists)
		assert.Equal(t, test.Expected, toArray(res))
		
		// check lists didn`t change
		assert.Equal(t, len(lists), len(listCopy))
		for i := 0; i < len(lists); i++ {
			assert.Equal(t, toArray(lists[i]), toArray(listCopy[i]))
		}
	}
}

func TestMergeListsInPlace(t *testing.T) {
	for _, test := range tests {
		lists := toArrayNodes(test.Lists)
		
		res := mergeKListsInPlace(lists)
		assert.Equal(t, test.Expected, toArray(res))
	}
}

func TestMergeListsWithCopy(t *testing.T) {
	for _, test := range tests {
		lists := toArrayNodes(test.Lists)
		listCopy := make([]*Node, len(lists))
		copy(listCopy, lists)
		
		res := mergeKListsWithCopy(lists)
		assert.Equal(t, test.Expected, toArray(res))
		
		// check lists didn`t change
		assert.Equal(t, len(lists), len(listCopy))
		for i := 0; i < len(lists); i++ {
			assert.Equal(t, toArray(lists[i]), toArray(listCopy[i]))
		}
	}
}