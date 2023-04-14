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

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

var max_r = 1

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	max_r = max(max_r, l+r+1) //比较了每一个点的最长路径，然后选出一个最大的
	return max(l, r) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	dfs(root)
	return max_r - 1
}

func main() {
	var head = new(TreeNode)
	head.val = 1

	var left = new(TreeNode)
	left.val = 2
	head.Left = left

	//var left1 = new(TreeNode)
	//left1.val = 4
	//left.Left = left1
	//
	//var left_right = new(TreeNode)
	//left_right.val = 5
	//left.Right = left_right
	//
	//var right = new(TreeNode)
	//right.val = 3
	//head.Right = right

	println(diameterOfBinaryTree(head))
}
