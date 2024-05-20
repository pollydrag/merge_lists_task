package main

///////////////////////////////////////////////////////////////
// InPlace solution
///////////////////////////////////////////////////////////////

func merge2ListsInPlace(a, b *Node) *Node {
	temp := &Node{}
	head := temp

	for a != nil && b != nil {
		if a.Value < b.Value {
			temp.Next = a
			a = a.Next
		} else {
			temp.Next = b
			b = b.Next
		}
		temp = temp.Next
	}

	if a != nil {
		temp.Next = a
	}

	if b != nil {
		temp.Next = b
	}

	return head.Next
}

func mergeKListsInPlace(lists []*Node) *Node {
	if len(lists) == 0 {
		return nil
	}
	
	for len(lists) > 1 {
		for i, j := 0, len(lists) - 1; i < j; i, j = i+1, j-1 {
			lists[i] = merge2ListsInPlace(lists[i], lists[j])
		}
		lists = lists[:(len(lists) + 1) / 2]
	}

	return lists[0]
}


///////////////////////////////////////////////////////////////
// Copy solution
///////////////////////////////////////////////////////////////
func merge2ListsWithCopy(a, b *Node) *Node {
	var root *Node
	prev := root

	for a != nil || b != nil {
		cur := &Node{}
		
		if root == nil {
			root = cur
		}
		
		if prev != nil {
			prev.Next = cur
		}
		
		prev = cur
		
		if a == nil {
			cur.Value = b.Value
			b = b.Next
			continue
		}
		if b == nil {
			cur.Value = a.Value
			a = a.Next
			continue
		}
		
		if a.Value < b.Value {
			cur.Value = a.Value
			a = a.Next
		} else {
			cur.Value = b.Value
			b = b.Next
		}
		
	}

	return root
}

func mergeKListsWithCopy(lists []*Node) *Node {
	if len(lists) == 0 {
		return nil
	}
	
	for len(lists) > 1 {
		newLists := make([]*Node, 0)
		for i, j := 0, len(lists) - 1; i <= j; i, j = i+1, j-1 {
			var res *Node
			if i == j {
				res = lists[i]
			} else {
				res = merge2ListsWithCopy(lists[i], lists[j])
			}
			newLists = append(newLists, res)
		}
		lists = newLists
	}

	return lists[0]
}

