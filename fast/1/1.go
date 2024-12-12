package __1

import "sort"

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	sum := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := target - nums[i]
		if _, ok := sum[nums[i]]; !ok {
			sum[n] = i
		} else {
			return []int{sum[nums[i]], i}
		}
	}
	return nil
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := &ListNode{}
	r := result
	b := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + b
		b = sum / 10
		r.Next = &ListNode{
			Val: sum % 10,
		}
		r = r.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil {
		for l2 != nil {
			sum := l2.Val + b
			b = sum / 10
			r.Next = &ListNode{
				Val: sum % 10,
			}
			r = r.Next
			l2 = l2.Next
		}
	}
	if l2 == nil {
		for l1 != nil {
			sum := l1.Val + b
			b = sum / 10
			r.Next = &ListNode{
				Val: sum % 10,
			}
			r = r.Next
			l1 = l1.Next
		}
	}
	if b == 1 {
		r.Next = &ListNode{
			Val: 1,
		}
	}
	return result.Next
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start

// 动态规划
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	maxLen := 1
	m := make(map[byte]bool)
	n := 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for n < len(s) && !m[s[n]] {
			m[s[n]] = true
			n++
		}
		if n-i > maxLen {
			maxLen = n - i
		}
	}
	return maxLen
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// 动态规划
// 突破口就是探测长度
// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	m := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		m[i] = make([]bool, len(s))
	}
	maxLen := 1
	start := 0
	for l := 2; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			j := i + l - 1
			if s[i] != s[j] {
				m[i][j] = false
				// 这里是3个的长度也算，存在中间一样的
			} else if l <= 3 {
				m[i][j] = true
			} else {
				m[i][j] = m[i+1][j-1]
			}
			if m[i][j] && l > maxLen {
				maxLen = l
				start = i
			}
		}
	}
	// 这里容易粗心，字符串分串是到最后一个之前
	return s[start : start+maxLen]
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 双指针
/*
最大水量=Math.min(height[left],height[right])*长度，往中间移动长度一定会变短

如果移动高的那一边，会有两种情况：

1、下一根柱子的高度比现在高，高度还取最小值低的那边，最大水量比原来小

2、下一根柱子的高度比现在低，高度比原来的最小值还小，最大水量比原来小

如果移动低的那一边，会有两种情况：

1、下一根柱子的高度比现在高，高度就可以取更高的值，最大水量不一定比原来小

2、下一根柱子的高度比现在低，高度比原来的最小值还小，最大水量比原来小

所以应该移动低的那一边
*/
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	l, r := 0, len(height)-1
	maxArea := 0
	for l < r {
		maxArea = max(maxArea, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return maxArea
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) (res [][]int) {
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j := len(nums) - 1
		target := -nums[i]
		for k := i + 1; k < j; k++ {
			if k != i+1 && nums[k] == nums[k-1] {
				continue
			}
			for k < j && nums[k]+nums[j] > target {
				j--
			}
			if j == k {
				break
			}
			if nums[k]+nums[j] == target {
				res = append(res, []int{nums[i], nums[k], nums[j]})
			}
		}
	}
	return res
}

// @lc code=end

var letterMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) (res []string) {
	if len(digits) == 0 {
		return nil
	}
	var fn func(int, string)
	fn = func(index int, s string) {
		if index == len(digits) {
			res = append(res, s)
			return
		}
		cur := letterMap[digits[index]]
		for i := 0; i < len(cur); i++ {
			// 取一个
			s += string(cur[i])
			// 找下一个
			fn(index+1, s)
			// 取消
			s = s[:len(s)-1]
		}
	}
	fn(0, "")
	return
}

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	// 其实是为了得到n+1个
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if cur == '(' && s[i] != ')' {
				return false
			} else if cur == '[' && s[i] != ']' {
				return false
			} else if cur == '{' && s[i] != '}' {
				return false
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	res := &ListNode{}
	temp := res
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			temp.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		} else {
			temp.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		}
		temp = temp.Next
	}
	if list1 == nil {
		temp.Next = list2
	}
	if list2 == nil {
		temp.Next = list1
	}
	return res.Next
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) (res []string) {
	var back func(left, right int, s string)
	back = func(left, right int, s string) {
		if len(s) == n*2 {
			res = append(res, s)
			return
		} else {
			if left < n {
				s += "("
				back(left+1, right, s)
				s = s[:len(s)-1]
			}
			if right < left {
				s += ")"
				back(left, right+1, s)
				s = s[:len(s)-1]
			}
		}
	}
	back(0, 0, "")
	return
}

// @lc code=end

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		// 正中
		if nums[mid] == target {
			return mid
		}
		// 先查看中间的值是否比最左边的值大
		if nums[mid] >= nums[0] {
			// 前面已经判断了不等于mid，所以不需要再写〈=
			if target >= nums[0] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			// 往左边找最小的
			curL := mid
			for curL-1 >= l {
				if nums[curL-1] == target {
					curL--
					continue
				}
				break
			}
			curR := mid
			for curR+1 <= r {
				if nums[curR+1] == target {
					curR++
					continue
				}
				break
			}
			return []int{curL, curR}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{-1, -1}
}

func searchInt(nums []int, target int, lower bool) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target || (lower && nums[mid] >= target) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// 左边是永远比target
	return l
}

func searchRange1(nums []int, target int) []int {
	l := searchInt(nums, target, true)
	r := searchInt(nums, target, false) - 1
	if l <= r && r < len(nums) && nums[l] == target && nums[r] == target {
		return []int{l, r}
	}
	return []int{-1, -1}
}

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) (res [][]int) {
	arr := make([]int, 0)
	var back func(cur, sum int)
	// cur是取零钱的类型
	back = func(cur, sum int) {
		if cur == len(candidates) {
			return
		}
		if sum == 0 {
			res = append(res, append([]int(nil), arr...))
			return
		}
		// 跳过这种零钱取下种
		back(cur+1, sum)
		if sum-candidates[cur] >= 0 {
			arr = append(arr, candidates[cur])
			back(cur, sum-candidates[cur])
			arr = arr[:len(arr)-1]
		}
	}
	back(0, target)
	return
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) (res [][]int) {
	var dfs func(int)
	dfs = func(curIndex int) {
		if curIndex == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}
		for i := curIndex; i < len(nums); i++ {
			nums[i], nums[curIndex] = nums[curIndex], nums[i]
			dfs(curIndex + 1)
			nums[i], nums[curIndex] = nums[curIndex], nums[i]
		}
	}
	dfs(0)
	return
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// 先对角线转 再平转
// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = matrix[j][i]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-i-1] = matrix[i][n-i-1], matrix[i][j]
		}
	}
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)
	for i := 0; i < len(strs); i++ {
		tmp := [26]int{}
		for j := 0; j < len(strs[i]); j++ {
			tmp[strs[i][j]-'a']++
		}
		res[tmp] = append(res[tmp], strs[i])
	}
	// r := make([][]string, 0, len(res))
	var r [][]string
	for _, v := range res {
		r = append(r, v)
	}
	return r
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 贪心算法 如果前面的结果为负数，无论如何相加都会小于本身
func maxSubArray1(nums []int) int {
	sum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		if sum > max {
			max = sum
		}
	}
	return max
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

func canJump(nums []int) bool {
	canJ := 0
	for i := 0; i < len(nums); i++ {
		if i <= canJ {
			canJ = max(canJ, i+nums[i])
		}
		if canJ >= len(nums)-1 {
			return true
		}
	}
	return false
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start

// 合并区间
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][]int) (res [][]int) {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > cur[1] {
			res = append(res, cur)
			cur = intervals[i]
		} else {
			cur[1] = max(cur[1], intervals[i][1])
		}
	}
	res = append(res, cur)
	return
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	row := len(grid)
	dp := make([][]int, row)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i-1 < 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j-1 < 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}
	return dp[row-1][col-1]
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	x, y, z := 0, 1, 1
	for i := 1; i < n; i++ {
		x, y = y, z
		z = x + y
	}
	return z
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start

func sortColors(nums []int) {
	l := 0
	r := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		for ; i < r && nums[i] == 2; r-- {
			nums[r], nums[i] = nums[i], nums[r]
		}
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) (res [][]int) {
	arr := make([]int, 0)
	var back func(cur int)
	back = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int(nil), arr...))
			return
		}
		// 添加一个
		arr = append(arr, nums[cur])
		back(cur + 1)
		// 撤销添加
		arr = arr[:len(arr)-1]
		// 添加一个空的
		back(cur + 1)
	}
	back(0)
	return
}

// @lc code=end

type pair struct {
	x, y int
}

var dirs = []pair{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	dp := make([][]bool, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]bool, col)
	}
	var check func(cur, x, y int) bool
	check = func(cur, x, y int) bool {
		if board[x][y] != word[cur] {
			return false
		}
		if cur == len(word)-1 {
			return true
		}
		dp[x][y] = true
		defer func() {
			dp[x][y] = false
		}()
		// 尝试往4个方向走
		for _, dir := range dirs {
			x1, y1 := x+dir.x, y+dir.y
			if x1 < 0 || x1 >= row {
				continue
			}
			if y1 < 0 || y1 >= col {
				continue
			}
			if dp[x1][y1] {
				// 走过了，就不要再走了
				continue
			}
			if check(cur+1, x1, y1) {
				return true
			}
		}
		return false
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if check(0, i, j) {
				return true
			}
		}
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var mid func(node *TreeNode)
	mid = func(node *TreeNode) {
		if node == nil {
			return
		}
		mid(node.Left)
		res = append(res, node.Val)
		mid(node.Right)
	}
	mid(root)
	return res
}

// @lc code=end
