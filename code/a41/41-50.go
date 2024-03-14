package a41

import (
	"math"
	"sort"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// Jump 45. 跳跃游戏2
// 贪心算法
// [2,3,1,1,4]
//
//	执行耗时:11 ms,击败了57.62% 的Go用户
//	内存消耗:5.9 MB,击败了88.53% 的Go用户
func Jump(nums []int) int {
	ans := 1
	cur, next := 0, 0
	curLen := len(nums)
	for next < curLen-1 {
		ans += 1
		tmp := next
		for i := cur; i < next+1; i++ {
			tmp = max(tmp, i+nums[i])
		}
		cur = next + 1
		next = tmp
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Jump1 跳跃游戏2
// 动态规划
// [2,3,1,1,4]
//
//	执行耗时:38 ms,击败了15.62% 的Go用户
//	内存消耗:5.9 MB,击败了19.11% 的Go用户
func Jump1(nums []int) int {
	curLen := len(nums)
	if curLen < 3 {
		return curLen - 1
	}
	dp := make([]int, curLen)
	for i := 0; i < curLen; i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 0; i < curLen-1; i++ {
		for j := 1; j < nums[i]+1; j++ {
			if i+j < curLen {
				dp[i+j] = min(dp[i+j], dp[i]+1)
			}
		}
	}
	return dp[curLen-1]
}

// Permute 46. 全排列
// 回溯
// [1,2,3,4,5]
//
//	执行耗时:2 ms,击败了26.62% 的Go用户
//	内存消耗:2.5 MB,击败了87.55% 的Go用户
func Permute(nums []int) [][]int {
	curLen := len(nums)
	var r [][]int
	var fn func(int, []int)
	fn = func(first int, output []int) {
		if first == curLen {
			tmp := make([]int, curLen)
			copy(tmp, output)
			r = append(r, tmp)
			return
		}
		for i := first; i < curLen; i++ {
			output[i], output[first] = output[first], output[i]
			fn(first+1, nums)
			// 撤销操作
			output[i], output[first] = output[first], output[i]
		}
	}
	fn(0, nums)
	return r
}

// Permute1 全排列(利用数组记录)
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.5 MB,击败了98.71% 的Go用户
func Permute1(nums []int) [][]int {
	curLen := len(nums)
	var r [][]int
	var fn func(int)
	var perm []int
	vis := make([]bool, curLen)
	fn = func(first int) {
		if first == curLen {
			tmp := make([]int, curLen)
			copy(tmp, perm)
			r = append(r, tmp)
			return
		}
		for i := 0; i < curLen; i++ {
			if vis[i] {
				continue
			}
			perm = append(perm, nums[i])
			vis[i] = true
			fn(first + 1)
			// 撤销操作
			perm = perm[:len(perm)-1]
			vis[i] = false
		}
	}
	fn(0)
	return r
}

// PermuteUnique 全排列2
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:3.5 MB,击败了85.79% 的Go用户
func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	curLen := len(nums)
	var r [][]int
	var fn func(int)
	var perm []int
	vis := make([]bool, curLen)
	fn = func(first int) {
		if first == curLen {
			tmp := make([]int, curLen)
			copy(tmp, perm)
			r = append(r, tmp)
			return
		}
		for i := 0; i < curLen; i++ {
			if vis[i] || i > 0 && !vis[i-1] && nums[i] == nums[i-1] {
				continue
			}
			perm = append(perm, nums[i])
			vis[i] = true
			fn(first + 1)
			// 撤销操作
			perm = perm[:len(perm)-1]
			vis[i] = false
		}
	}
	fn(0)
	return r
}

// Rotate 旋转图像
// 先上下翻转 再对角线翻转
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.1 MB,击败了77.70% 的Go用户
func Rotate(matrix [][]int) {
	curLen := len(matrix)
	for i := 0; i < curLen/2; i++ {
		for j := 0; j < curLen; j++ {
			matrix[i][j], matrix[curLen-i-1][j] = matrix[curLen-i-1][j], matrix[i][j]
		}
	}
	for i := 0; i < curLen; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// GroupAnagrams 49. 字母异位词分组
//
//	执行耗时:22 ms,击败了32.59% 的Go用户
//	内存消耗:7.6 MB,击败了65.58% 的Go用户
func GroupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		v := []byte(strs[i])
		sort.Slice(v, func(a, b int) bool {
			return v[a] > v[b]
		})
		s := string(v)
		arr, ok := m[s]
		if !ok {
			arr = make([]string, 0)
		}
		arr = append(arr, strs[i])
		m[s] = arr
	}
	var r [][]string
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// MyPow 50. Pow(x, n)
// 一个一个乘会超时，所以需要优化算法
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:1.9 MB,击败了31.46% 的Go用户
func MyPow(x float64, n int) float64 {
	if n >= 0 {
		return computeMyPow(x, n)
	}
	return 1.0 / computeMyPow(x, n)
}

func computeMyPow(v float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := computeMyPow(v, n/2)
	if n%2 == 0 {
		return y * y
	}
	return v * y * y
}
