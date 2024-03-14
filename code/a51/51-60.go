package a51

import (
	"sort"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MaxSubArray 最大子数组和
// [-2,1,-3,4,-1,2,1,-5,4]
// 贪心算法
//
//	执行耗时:81 ms,击败了62.69% 的Go用户
//	内存消耗:8.3 MB,击败了43.19% 的Go用户
func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ans := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ans = max(ans, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

// MaxSubArray1 最大子数组和
// [-2,1,-3,4,-1,2,1,-5,4]
// 动态规划
// d[n] = max(d[n-1] + num[n], num[n])
//
//	执行耗时:83 ms,击败了58.47% 的Go用户
//	内存消耗:8.1 MB,击败了63.87% 的Go用户
func MaxSubArray1(nums []int) int {
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i-1]+nums[i], nums[i])
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	return maxNum
}

// SpiralOrder 螺旋矩阵
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:1.9 MB,击败了86.93% 的Go用户
func SpiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 1 {
		return matrix[0]
	}
	n := len(matrix[0])
	left, right, top, bottom := 0, n-1, 0, m-1
	var r []int
	for {
		for i := left; i <= right; i++ {
			r = append(r, matrix[top][i])
		}
		top++
		if left > right {
			break
		}

		for i := top; i <= bottom; i++ {
			r = append(r, matrix[i][right])
		}
		right--
		if top > bottom {
			break
		}

		for i := right; i >= left; i-- {
			r = append(r, matrix[bottom][i])
		}
		bottom--
		if left > right {
			break
		}

		for i := bottom; i >= top; i-- {
			r = append(r, matrix[i][left])
		}
		left++
		if top > bottom {
			break
		}
	}
	return r
}

// CanJump 跳跃游戏
// 贪心算法
//
//	执行耗时:42 ms,击败了82.90% 的Go用户
//	内存消耗:6.7 MB,击败了83.21% 的Go用户
func CanJump(nums []int) bool {
	cur, next := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		tmp := next
		for j := cur; j < len(nums)-1 && j < next+1; j++ {
			tmp = max(tmp, j+nums[j])
		}
		cur = next + 1
		next = tmp
	}
	return next >= len(nums)-1
}

// CanJump1
// 从后往前跳
// [2,3,1,1,4]
// [3,2,1,0,4]
//
//	执行耗时:44 ms,击败了77.52% 的Go用户
//	内存消耗:6.8 MB,击败了44.08% 的Go用户
func CanJump1(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	curLen := len(nums)
	minNext := curLen - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= minNext {
			minNext = i
		}
	}
	return minNext == 0
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Merge 56. 合并区间
// [[1,3],[2,6],[8,10],[15,18]]
//
//	执行耗时:16 ms,击败了75.39% 的Go用户
//	内存消耗:6.1 MB,击败了53.04% 的Go用户
func Merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	judge := func(a, b []int) bool {
		return a[1] >= b[0] && b[1] >= a[0]
	}
	ans := [][]int{
		intervals[0],
	}
	for i := 1; i < len(intervals); i++ {
		num := len(ans) - 1
		if judge(ans[num], intervals[i]) {
			ans[num][0] = min(ans[num][0], intervals[i][0])
			ans[num][1] = max(ans[num][1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

// Insert 57. 插入区间
//
//	执行耗时:4 ms,击败了85.07% 的Go用户
//	内存消耗:4.2 MB,击败了20.06% 的Go用户
func Insert(intervals [][]int, newInterval []int) [][]int {
	if newInterval == nil {
		return intervals
	}
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	judge := func(a, b []int) bool {
		return a[1] >= b[0] && b[1] >= a[0]
	}
	ans := [][]int{
		intervals[0],
	}
	for i := 1; i < len(intervals); i++ {
		num := len(ans) - 1
		if judge(ans[num], intervals[i]) {
			ans[num][0] = min(ans[num][0], intervals[i][0])
			ans[num][1] = max(ans[num][1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

// LengthOfLastWord 58. 最后一个单词的长度
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2 MB,击败了42.56% 的Go用户
func LengthOfLastWord(s string) int {
	num := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if num > 0 {
				return num
			}
			continue
		}
		num++
	}
	return num
}

// GenerateMatrix 59. 螺旋矩阵II
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2 MB,击败了96.11% 的Go用户
func GenerateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	num := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for {
		for i := left; i <= right; i++ {
			ans[top][i] = num
			num++
		}
		top++
		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			ans[i][right] = num
			num++
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			ans[bottom][i] = num
			num++
		}
		bottom--
		if top > bottom {
			break
		}

		for i := bottom; i >= top; i-- {
			ans[i][left] = num
			num++
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}
