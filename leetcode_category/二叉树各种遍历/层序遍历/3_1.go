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

// levelOrder 实现了二叉树的层次遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int            // 用于存储结果
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
		res = append(res, level) // 将当前层的节点值加入到结果中
	}
	return res
}

// 算法的步骤如下：
//
// 1.创建一个空队列，将根节点入队。
// 2.当队列不为空时，重复步骤 3 到 5，否则遍历结束。
// 3.将队首节点出队，访问该节点。
// 4.将该节点的左儿子入队（如果有的话）。
// 5.将该节点的右儿子入队（如果有的话）。
func main() {
	println(levelOrder(createTree()))
}
