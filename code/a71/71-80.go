package a71

import (
	"strings"
)

// SimplifyPath 71. 简化路径
//
//	执行耗时:2 ms,击败了37.93% 的Go用户
//	内存消耗:3 MB,击败了48.58% 的Go用户
func SimplifyPath(path string) string {
	var stack []string
	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

// MinDistance 72. 编辑距离
// 执行耗时:3 ms,击败了64.20% 的Go用户
// 内存消耗:5.3 MB,击败了69.72% 的Go用户
func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[m][n]
}

func min(x, y, z int) int {
	if x > y {
		if y < z {
			return y
		} else {
			return z
		}
	}
	if x < z {
		return x
	}
	return z
}

// SetZeroes 73. 矩阵置零
// 用两个数组 遍历两遍
//
//	执行耗时:8 ms,击败了78.15% 的Go用户
//	内存消耗:5.5 MB,击败了81.90% 的Go用户
func SetZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// SetZeroes1 73. 矩阵置零
// 使用第一行和一个列作为额外数组
//
//	执行耗时:10 ms,击败了42.02% 的Go用户
//	内存消耗:5.5 MB,击败了59.82% 的Go用户
func SetZeroes1(matrix [][]int) {
	firstRow, firstCol := false, false
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			firstRow = true
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstCol = true
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if firstRow {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if firstCol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

// SearchMatrix 74. 搜索二纬矩阵
// 二分查找
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.4 MB,击败了29.08% 的Go用户
func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		row := mid / n
		col := mid % n
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// SortColors 75. 颜色分类
// [2,0,2,1,1,0]
// 单指针 遍历两遍
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2 MB,击败了74.29% 的Go用户
func SortColors(nums []int) {
	n := swapNumber(nums, 0)
	swapNumber(nums[n:], 1)
}

func swapNumber(nums []int, target int) int {
	x := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			nums[x], nums[i] = nums[i], nums[x]
			x++
		}
	}
	return x
}

// SortColors1 75. 颜色分类
// [2,0,2,1,1,0]
// 双指针 都从0开始
//
//	执行耗时:1 ms,击败了23.45% 的Go用户
//	内存消耗:2 MB,击败了97.18% 的Go用户
func SortColors1(nums []int) {
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

// SortColors2 75. 颜色分类
// [2,0,2,1,1,0]
// [2,0,1]
// 双指针 p0和p2 更加好理解
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2 MB,击败了97.18% 的Go用户
func SortColors2(nums []int) {
	if len(nums) < 2 {
		return
	}
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for i <= p2 && nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
		}
	}
}

// Combine 77. 组合
// 难点在于如何去重复，发现规律右边一定会小于右边，这样我们下个循环就只取当前大于i+1的
//
//	执行耗时:137 ms,击败了36.24% 的Go用户
//	内存消耗:63.4 MB,击败了42.24% 的Go用户
func Combine(n int, k int) [][]int {
	var fn func(int, int)
	var r [][]int
	var ans []int
	fn = func(j int, cur int) {
		if cur == k {
			tmp := make([]int, k)
			copy(tmp, ans)
			r = append(r, tmp)
			return
		}
		for i := j; i <= n; i++ {
			ans = append(ans, i)
			fn(i+1, cur+1)
			ans = ans[:len(ans)-1]
		}
	}
	fn(1, 0)
	return r
}

// Subsets 78. 子集
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.2 MB,击败了84.95% 的Go用户
func Subsets(nums []int) [][]int {
	var fn func(int, int)
	var r [][]int
	var ans []int
	curLen := len(nums)
	fn = func(j int, cur int) {
		tmp := make([]int, len(ans))
		copy(tmp, ans)
		r = append(r, tmp)
		if cur == curLen {
			return
		}
		for i := j; i < len(nums); i++ {
			ans = append(ans, nums[i])
			fn(i+1, cur+1)
			ans = ans[:len(ans)-1]
		}
	}
	fn(0, 0)
	return r
}

// Exist 单词搜索
//
//	执行耗时:76 ms,击败了97.09% 的Go用户
//	内存消耗:1.9 MB,击败了76.55% 的Go用户
func Exist(board [][]byte, word string) bool {
	m, n, wN := len(board), len(board[0]), len(word)
	var findNext func(int, int, int, string) bool
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	findNext = func(x, y, cur int, dir string) bool {
		if vis[x][y] {
			return false
		}
		if cur == wN {
			return true
		}
		// 看看左边
		if dir != "right" && y-1 >= 0 && board[x][y-1] == word[cur] {
			vis[x][y] = true
			if findNext(x, y-1, cur+1, "left") {
				return true
			}
			vis[x][y] = false
		}
		// 右边
		if dir != "left" && y+1 < n && board[x][y+1] == word[cur] {
			vis[x][y] = true
			if findNext(x, y+1, cur+1, "right") {
				return true
			}
			vis[x][y] = false
		}
		// 上边
		if dir != "bottom" && x-1 >= 0 && board[x-1][y] == word[cur] {
			vis[x][y] = true
			if findNext(x-1, y, cur+1, "top") {
				return true
			}
			vis[x][y] = false
		}
		// 下边
		if dir != "top" && x+1 < m && board[x+1][y] == word[cur] {
			vis[x][y] = true
			if findNext(x+1, y, cur+1, "bottom") {
				return true
			}
			vis[x][y] = false
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != word[0] {
				continue
			}
			// 从这个点开始找
			if findNext(i, j, 1, "") {
				return true
			}
		}
	}
	return false
}

// RemoveDuplicates 删除有序数组的重复项II
// [1,1,1,2,2,3]
//
//	执行耗时:3 ms,击败了62.30% 的Go用户
//	内存消耗:2.7 MB,击败了29.15% 的Go用户
func RemoveDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	left := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[left-1] && nums[i] == nums[left-2] {
			continue
		}
		nums[i], nums[left] = nums[left], nums[i]
		left++
	}
	return left
}
