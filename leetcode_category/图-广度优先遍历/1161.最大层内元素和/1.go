package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func maxLevelSum(root *TreeNode) int {
	/**
	 * Definition for a binary tree node.
	 * type TreeNode struct {
	 *     Val int
	 *     Left *TreeNode
	 *     Right *TreeNode
	 * }
	 */
	max := root.Val
	max_l, res := 1, 0
	// var res [][]int            // 用于存储结果
	queue := []*TreeNode{root} // 创建一个队列，将根节点入队
	for len(queue) > 0 {       // 当队列不为空时，重复以下步骤
		var level []int                   // 用于存储当前层的节点值
		for i := len(queue); i > 0; i-- { // 遍历当前层的所有节点
			node := queue[0]                // 取出队首节点
			queue = queue[1:]               // 将队首节点出队（队在这里减小）
			level = append(level, node.Val) // 将节点值加入到 level 中
			if node.Left != nil {           // 如果左子节点不为空，则将其入队
				queue = append(queue, node.Left)
			}
			if node.Right != nil { // 如果右子节点不为空，则将其入队
				queue = append(queue, node.Right)
			}
		}
		res += 1
		sum := 0
		for _, x := range level {
			sum += x
		}
		if sum > max {
			max = sum
			max_l = res
		}
	}
	return max_l
}
