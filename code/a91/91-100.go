package a91

import (
	"math"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 91. 解码方法
// 动态规划
// 两种情况
// 1. 当单独拆解的时候，当s[i] != "0"的时候，处理方式会和f[i-1]一样
// 2. 当两个字符串一起的时候。如果s[i-1] != "0"，并且两个字符串组合在一块不超过26，处理方式和f[i-2]一样
// 此时f[i] = f[i-1] + f[i-2]
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:1.9 MB,击败了81.73% 的Go用户
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && int(s[i-2]-'0')*10+int(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

// ReverseList 反转链表
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.4 MB,击败了44.01% 的Go用户
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// ReverseBetween 92. 反转链表II
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2 MB,击败了77.68% 的Go用户
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	// 先找到头前一个
	for i := 0; i < left-1; i++ {
		cur = cur.Next
	}
	head2 := cur.Next
	head1 := cur
	head1.Next = nil
	head3 := head2
	for i := 0; i < right-left; i++ {
		head3 = head3.Next
	}
	// 如果h3为nil
	var end *ListNode
	if head3.Next != nil {
		end = head3.Next
	}
	head3.Next = nil
	ReverseList(head2)
	head1.Next = head3
	head2.Next = end
	return dummy.Next
}

// RestoreIpAddresses 93. 复原IP地址
// "25525511135"
// 回溯 + 剪枝
//
//	执行耗时:1 ms,击败了26.97% 的Go用户
//	内存消耗:2 MB,击败了69.60% 的Go用户
func RestoreIpAddresses(s string) []string {
	var r []string
	var ans []string
	var fn func(int, int)
	fn = func(i int, n int) {
		if n == 4 && i == len(s) {
			r = append(r, strings.Join(ans, "."))
			return
		}
		if n == 4 || i == len(s) {
			return
		}
		if s[i] == '0' {
			ans = append(ans, "0")
			fn(i+1, n+1)
			ans = ans[:len(ans)-1]
			return
		}
		for l := 1; l <= 3 && i+l <= len(s); l++ {
			if l == 3 {
				cur, _ := strconv.Atoi(s[i : i+l])
				if cur > 255 {
					continue
				}
			}
			ans = append(ans, s[i:i+l])
			fn(i+l, n+1)
			ans = ans[:len(ans)-1]
		}
	}
	fn(0, 0)
	return r
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderTraversal 二叉树的中序遍历
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:1.9 MB,击败了23.84% 的Go用户
func inorderTraversal(root *TreeNode) []int {
	var r []int
	var fn func(node *TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		fn(node.Left)
		r = append(r, node.Val)
		fn(node.Right)
	}
	fn(root)
	return r
}

// GenerateTrees 不同的二叉搜索数II
// 递归
//
//	执行耗时:3 ms,击败了45.45% 的Go用户
//	内存消耗:4.2 MB,击败了80.68% 的Go用户
func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var fn func(start int, end int) []*TreeNode
	fn = func(start int, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		var d []*TreeNode
		for i := start; i <= end; i++ {
			leftArr := fn(start, i-1)
			rightArr := fn(i+1, end)
			for j := 0; j < len(leftArr); j++ {
				for k := 0; k < len(rightArr); k++ {
					cur := &TreeNode{Val: i, Left: nil, Right: nil}
					cur.Left = leftArr[j]
					cur.Right = rightArr[k]
					d = append(d, cur)
				}
			}
		}
		return d
	}
	return fn(1, n)
}

// NumTrees 不同的二叉搜索数
// dp[i] = d[i, 1] + d[i, 2] + ... d[i, i]
// d[i, 1]表示i个数，以1为根的二叉搜索数的数量
// 左边是j-1个数，右边是i-j个数
// 可以登出d[i, j] = d[j-1] * d[i - j]
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:1.8 MB,击败了34.25% 的Go用户
func NumTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		sum := 0
		for j := 1; j <= i; j++ {
			sum += dp[j-1] * dp[i-j]
		}
		dp[i] = sum
	}
	return dp[n]
}

// isValidBST 验证二叉搜索数
//
//	执行耗时:5 ms,击败了35.83% 的Go用户
//	内存消耗:5.1 MB,击败了87.06% 的Go用户
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return IsValidBSTFn(root, math.MinInt64, math.MaxInt64)
}

func IsValidBSTFn(node *TreeNode, miNum, maxNum int) bool {
	if node == nil {
		return true
	}
	return node.Val > miNum && node.Val < maxNum && IsValidBSTFn(node.Left, miNum, node.Val) && IsValidBSTFn(node.Right, node.Val, maxNum)
}

// 中序遍历
//
//	执行耗时:5 ms,击败了35.83% 的Go用户
//	内存消耗:5.2 MB,击败了27.12% 的Go用户
func isValidBST1(root *TreeNode) bool {
	minNum := math.MinInt64
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= minNum {
			return false
		}
		minNum = root.Val
		root = root.Right
	}
	return true
}

// recoverTree 恢复二叉搜索数
// 解法1. 将数中序遍历得到一个列表，然后找到两个破坏有序的数，交换位置
// 双指针需要考虑两个值先连的情况
//
//	执行耗时:5 ms,击败了76.61% 的Go用户
//	内存消耗:5.7 MB,击败了35.32% 的Go用户
func recoverTree(root *TreeNode) {
	var r []*TreeNode
	var fn func(node *TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		fn(node.Left)
		r = append(r, node)
		fn(node.Right)
	}
	fn(root)

	var x, y *TreeNode
	for i := 0; i < len(r)-1; i++ {
		if r[i+1].Val < r[i].Val {
			if x == nil {
				x, y = r[i], r[i+1]
			} else {
				y = r[i+1]
			}
		}
	}
	if x != nil && y != nil {
		x.Val, y.Val = y.Val, x.Val
	}
}

// recoverTree 恢复二叉搜索数
// 解法2. 用pre来记录上一个值
// 双指针需要考虑两个值先连的情况
//
//	执行耗时:13 ms,击败了13.30% 的Go用户
//	内存消耗:5.3 MB,击败了97.25% 的Go用户
func recoverTree1(root *TreeNode) {
	var fn func(node *TreeNode)
	var pre, x, y *TreeNode
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		fn(node.Left)
		if pre == nil {
			pre = node
		} else {
			if pre.Val > node.Val {
				y = node
				if x == nil {
					x = pre
				}
			}
			pre = node
		}
		fn(node.Right)
	}
	fn(root)
	if x != nil && y != nil {
		x.Val, y.Val = y.Val, x.Val
	}
}

//  100. 相同的数
//     执行耗时:0 ms,击败了100.00% 的Go用户
//     内存消耗:2 MB,击败了78.08% 的Go用户
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
