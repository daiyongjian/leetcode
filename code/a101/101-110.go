package a101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 对称二叉树
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.6 MB,击败了79.56% 的Go用户
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricFn(root.Left, root.Right)
}

func isSymmetricFn(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isSymmetricFn(left.Left, right.Right) && isSymmetricFn(left.Right, right.Left)
}

// LevelOrder 二叉树的层序遍历
//
//	执行耗时:2 ms,击败了49.88% 的Go用户
//	内存消耗:3.5 MB,击败了21.35% 的Go用户
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var r [][]int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		n := len(stack)
		var d []int
		for i := 0; i < n; i++ {
			cur := stack[i]
			if cur == nil {
				continue
			}
			d = append(d, cur.Val)
			stack = append(stack, cur.Left)
			stack = append(stack, cur.Right)
		}
		stack = stack[n:]
		if d == nil {
			continue
		}
		r = append(r, d)
	}
	return r
}

// ZigzagLevelOrder 二叉树的锯齿形层序遍历
// 也可以获取正常的队列后，将队列翻转
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:5.5 MB,击败了5.27% 的Go用户
func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var r [][]int
	var stack []*TreeNode
	stack = append(stack, root)
	dir := true
	for len(stack) != 0 {
		n := len(stack)
		var d []int
		add := 0
		for i := 0; i < n; i++ {
			var cur *TreeNode
			if dir {
				// 从左往右
				cur = stack[i]
			} else {
				// 从右向左
				cur = stack[n+add-i-1]
			}
			if cur == nil {
				continue
			}
			add += 2
			d = append(d, cur.Val)
			if dir {
				stack = append(stack, cur.Left)
				stack = append(stack, cur.Right)
			} else {
				stack = append([]*TreeNode{cur.Right}, stack...)
				stack = append([]*TreeNode{cur.Left}, stack...)
			}
		}
		if dir {
			stack = stack[n:]
		} else {
			stack = stack[:len(stack)-n]
		}
		dir = !dir
		if d == nil {
			continue
		}
		r = append(r, d)
	}
	return r
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// maxDepth 二叉树最大深度
// 深度优先
//
//	执行耗时:4 ms,击败了48.18% 的Go用户
//	内存消耗:4.2 MB,击败了98.48% 的Go用户
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// MaxDepth1 广度优先
//
//	执行耗时:3 ms,击败了61.25% 的Go用户
//	内存消耗:4.3 MB,击败了13.01% 的Go用户
func MaxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var stack []*TreeNode
	stack = append(stack, root)
	level := 0
	for ; len(stack) > 0; level++ {
		n := len(stack)
		for i := 0; i < n; i++ {
			cur := stack[i]
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
		}
		stack = stack[n:]
	}
	return level
}

// buildTree1 从前序与中序遍历序列构建二叉树
//
//	执行耗时:3 ms,击败了84.72% 的Go用户
//	内存消耗:3.9 MB,击败了72.29% 的Go用户
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = buildTree1(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree1(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// buildTree 从中序与后序遍历序列构建二叉树
//
//	执行耗时:6 ms,击败了27.35% 的Go用户
//	内存消耗:4 MB,击败了29.90% 的Go用户
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	inMap := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		inMap[inorder[i]] = i
	}
	var fn func(int, int) *TreeNode
	fn = func(start int, end int) *TreeNode {
		if start > end {
			return nil
		}
		root := &TreeNode{
			Val: postorder[len(postorder)-1],
		}
		postorder = postorder[:len(postorder)-1]
		i := inMap[root.Val]
		root.Right = fn(i+1, end)
		root.Left = fn(start, i-1)
		return root
	}
	return fn(0, len(inorder)-1)
}

// levelOrderBottom 二叉树的层序遍历II
//
//	执行耗时:1 ms,击败了38.47% 的Go用户
//	内存消耗:2.7 MB,击败了44.99% 的Go用户
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var r [][]int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		var m []int
		n := len(stack)
		for i := 0; i < n; i++ {
			cur := stack[i]
			m = append(m, cur.Val)
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
		}
		stack = stack[n:]
		r = append(r, m)
	}
	// 交换一下列表
	curLen := len(r)
	for i := 0; i < curLen/2; i++ {
		r[i], r[curLen-i-1] = r[curLen-i-1], r[i]
	}
	return r
}

// sortedArrayToBST 108. 将有序数组转化为二叉搜索数
//
//	执行耗时:2 ms,击败了46.53% 的Go用户
//	内存消耗:3.3 MB,击败了15.85% 的Go用户
func sortedArrayToBST(nums []int) *TreeNode {
	var fn func(left, right int) *TreeNode
	fn = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = fn(left, mid-1)
		root.Right = fn(mid+1, right)
		return root
	}
	return fn(0, len(nums)-1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// sortedListToBST 将有序链表转换为二叉搜索数
// 跟链表一样的做法
//
//	执行耗时:9 ms,击败了21.93% 的Go用户
//	内存消耗:5.6 MB,击败了86.62% 的Go用户
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{
		Val: slow.Val,
	}
	if pre != nil {
		pre.Next = nil
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(slow.Next)
	return root
}

// isBalanced 是否是平衡二叉树
//
//	执行耗时:3 ms,击败了88.04% 的Go用户
//	内存消耗:5.4 MB,击败了100.00% 的Go用户
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	n := maxDepth(root.Left) - maxDepth(root.Right)
	if n < 0 {
		n = -n
	}
	return n <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
