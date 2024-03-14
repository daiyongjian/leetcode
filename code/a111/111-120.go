package a111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// minDepth 二叉树的最小深度
// 注意有单边的存在
// 执行耗时:161 ms,击败了42.78% 的Go用户
// 内存消耗:21.9 MB,击败了15.42% 的Go用户
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// hasPathSum 路径总和
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:4.5 MB,击败了99.58% 的Go用户
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 这个才是叶子节点
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 执行耗时:3 ms,击败了66.43% 的Go用户
// 内存消耗:4.2 MB,击败了69.22% 的Go用户
func pathSum(root *TreeNode, targetSum int) [][]int {
	var r [][]int
	var ans []int
	var fn func(node *TreeNode, n int)
	fn = func(node *TreeNode, n int) {
		if node == nil {
			return
		}
		n -= node.Val
		ans = append(ans, node.Val)
		if node.Left == nil && node.Right == nil {
			if n == 0 {
				tmp := make([]int, len(ans))
				copy(tmp, ans)
				r = append(r, tmp)
			}
			ans = ans[:len(ans)-1]
			return
		}
		fn(node.Left, n)
		fn(node.Right, n)
		ans = ans[:len(ans)-1]
	}
	fn(root, targetSum)
	return r
}

// 用到了栈 效率很低
//
//	执行耗时:2 ms,击败了37.36% 的Go用户
//	内存消耗:2.8 MB,击败了27.49% 的Go用户
func flatten1(root *TreeNode) {
	var stack []*TreeNode
	var fn func(node *TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		stack = append(stack, node)
		fn(node.Left)
		fn(node.Right)
	}
	fn(root)
	cur := root
	for i := 0; i < len(stack)-1; i++ {
		cur.Left = nil
		cur.Right = stack[i+1]
		cur = cur.Right
	}
}

// 想办法原地旋转
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.7 MB,击败了99.65% 的Go用户
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		tmp := root.Right
		root.Right = root.Left
		root.Left = nil
		for root.Right != nil {
			root = root.Right
		}
		root.Right = tmp
	}
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 填充每个节点的下一个右侧节点指针
//
//	执行耗时:4 ms,击败了71.50% 的Go用户
//	内存消耗:6.3 MB,击败了49.22% 的Go用户
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var stack []*Node
	stack = append(stack, root)
	for len(stack) > 0 {
		n := len(stack)
		for i := 0; i < n; i++ {
			cur := stack[i]
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if i < n-1 {
				cur.Next = stack[i+1]
			}
		}
		stack = stack[n:]
	}
	return root
}

// Generate 杨辉三角
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.2 MB,击败了66.10% 的Go用
func Generate(numRows int) [][]int {
	r := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		r[i] = make([]int, i+1)
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				r[i][j] = 1
			} else if j == i {
				r[i][j] = 1
			} else {
				r[i][j] = r[i-1][j-1] + r[i-1][j]
			}
		}
	}
	return r
}

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.9 MB,击败了80.78% 的Go用户
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}

func minimumTotal(triangle [][]int) int {
}
