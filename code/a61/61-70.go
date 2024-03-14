package a61

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// RotateRight 61. 旋转链表
// 思路：先找到最后k个链表位置，然后将该位置指回头部
// 难点 k可能会超过链表长度
// 效率太低
//
//	执行耗时:848 ms,击败了5.00% 的Go用户
//	内存消耗:2.4 MB,击败了9.66% 的Go用户
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	fake := &ListNode{Next: head}
	first := fake
	end := fake
	for i := 0; i < k+1; i++ {
		if end.Next == nil {
			end = fake
		}
		end = end.Next
	}
	for end != nil {
		first = first.Next
		end = end.Next
	}
	top := first.Next
	if top == nil {
		return fake.Next
	}
	next := top
	first.Next = nil
	// 直到最后一个
	for next.Next != nil {
		next = next.Next
	}
	next.Next = fake.Next
	return top
}

// RotateRight1 61. 旋转链表
// 思路：将链表变成环
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.4 MB,击败了34.77% 的Go用户
func RotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
		n++
	}
	// 这里是倒数第N个
	k1 := n - k%n
	if n == k {
		return head
	}
	tmp.Next = head
	for i := 0; i < k1; i++ {
		tmp = tmp.Next
	}
	top := tmp.Next
	tmp.Next = nil
	return top
}

// UniquePaths 不同路径
// 动态规划
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:1.9 MB,击败了78.75% 的Go用户
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	if m == 1 || n == 1 {
		return 1
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else if i == 0 {
				dp[0][j] = dp[0][j-1]
			} else {
				dp[i][0] = dp[i-1][0]
			}
		}
	}
	return dp[m-1][n-1]
}

// UniquePathsWithObstacles 不同路径II
// 动态规划
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.3 MB,击败了75.23% 的Go用户
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else if i == 0 {
				dp[0][j] = dp[0][j-1]
			} else {
				dp[i][0] = dp[i-1][0]
			}
		}
	}
	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// MinPathSum 最小路径和
//
//	执行耗时:5 ms,击败了43.61% 的Go用户
//	内存消耗:3.7 MB,击败了58.48% 的Go用户
func MinPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if i > 0 && j > 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			} else if i == 0 {
				dp[0][j] = dp[0][j-1] + grid[i][j]
			} else {
				dp[i][0] = dp[i-1][0] + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

// plusOne 加一
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:1.9 MB,击败了33.37% 的Go用户
func plusOne(digits []int) []int {
	level := 0
	for i := len(digits) - 1; i >= 0; i-- {
		var sum int
		if i == len(digits)-1 {
			sum = digits[i] + level + 1
		} else {
			sum = digits[i] + level
		}
		level = sum / 10
		digits[i] = sum % 10
	}
	if level == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// AddBinary 二进制求和
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:3.4 MB,击败了51.67% 的Go用户
func AddBinary(a string, b string) string {
	d := ""
	level := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := level
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		level = sum / 2
		d = strconv.Itoa(sum%2) + d
	}
	if level == 1 {
		d = "1" + d
	}
	return d
}

// MySqrt x的平方根
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2 MB,击败了32.35% 的Go用户
func MySqrt(x int) int {
	left, right := 0, x
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		s := mid * mid
		if s == x {
			return mid
		} else if s > x {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}

// ClimbStairs 爬楼梯
// 递归
func ClimbStairs(n int) int {
	var fn func(int) int
	fn = func(i int) int {
		if i == 0 {
			return 0
		} else if i == 1 {
			return 1
		} else if i == 2 {
			return 2
		} else {
			return fn(i-1) + fn(i-2)
		}
	}
	return fn(n)
}

// ClimbStairs1 爬楼梯
// 迭代
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:1.8 MB,击败了60.57% 的Go用户
func ClimbStairs1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	x, y, z := 1, 2, 0
	for i := 2; i < n; i++ {
		z = x + y
		x, y = y, z
	}
	return z
}
