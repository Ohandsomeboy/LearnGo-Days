package main

//对于子树到父数的直径最大值有三种可能
//1.取左子树直径最大值
//2.取右子树直径最大值
//3.取经过当前节点最大值，左经过根最大长度+右经过根最大长度+2
//
//因此我们直到有两个因素是影响答案的重要值：
//1.以当前节点为根节点的最大值
//2.经过当前节点或子树节点的最大直径

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

// max返回两个整数中的最大值
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

var max_r = 1

// dfs是一个辅助函数，用于遍历二叉树并返回其子树的最大长度
func dfs(root *TreeNode, max_r *int) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left, max_r)
	r := dfs(root.Right, max_r)
	*max_r = max(*max_r, l+r+1) // 比较左右子树的最大长度
	return max(l, r) + 1        // 返回左右子树的最大长度加1
}

// diameterOfBinaryTree是一个计算二叉树直径的函数
func diameterOfBinaryTree(root *TreeNode) int {
	var max_r = 1     // 将max_r初始化为1
	dfs(root, &max_r) // 遍历二叉树并使用其直径更新max_r
	return max_r - 1  // 返回二叉树的直径
}

func createTree() *TreeNode {
	root := &TreeNode{val: 1}
	root.Left = &TreeNode{val: 2}
	root.Right = &TreeNode{val: 3}
	root.Left.Left = &TreeNode{val: 4}
	root.Left.Right = &TreeNode{val: 5}
	root.Right.Left = &TreeNode{val: 6}
	root.Right.Right = &TreeNode{val: 7}
	return root
}

func main() {
	println(diameterOfBinaryTree(createTree()))
}
