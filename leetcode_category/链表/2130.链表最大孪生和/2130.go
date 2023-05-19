package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	nums := make([]int, 0, 10000)
	p := &ListNode{Next: head} // 和head一个指向，空节点
	c := 0
	mn := 0
	d := p.Next
	for d != nil {
		nums = append(nums, d.Val)
		d = d.Next
		c++
	}
	for i := 0; i < c/2; i++ {
		if nums[i]+nums[c-1-i] > mn {
			mn = nums[i] + nums[c-i-1]
		}
	}
	return mn
}

func main() {
	var head = new(ListNode)
	head.Val = 5
	var Node1 = new(ListNode)
	head.Next = Node1
	Node1.Val = 4
	var Node2 = new(ListNode)
	Node1.Next = Node2
	Node2.Val = 2
	var Node3 = new(ListNode)
	Node2.Next = Node3
	Node3.Val = 1
	fmt.Println(pairSum(head))
}
