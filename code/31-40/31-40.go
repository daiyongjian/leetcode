package _1_40

import (
	"fmt"
	"sort"
)

// NextPermutation 下一个排列
// 这不是人能想的到的
// 我们希望下一个数 比当前数大，这样才满足 “下一个排列” 的定义。因此只需要 将后面的「大数」与前面的「小数」交换，就能得到一个更大的数。比如 123456，将 5 和 6 交换就能得到一个更大的数 123465。
// 我们还希望下一个数 增加的幅度尽可能的小，这样才满足“下一个排列与当前排列紧邻“的要求。为了满足这个要求，我们需要：
// 在 尽可能靠右的低位 进行交换，需要 从后向前 查找
// 将一个 尽可能小的「大数」 与前面的「小数」交换。比如 123465，下一个排列应该把 5 和 4 交换而不是把 6 和 4 交换
// 将「大数」换到前面后，需要将「大数」后面的所有数 重置为升序，升序排列就是最小的排列。以 123465 为例：首先按照上一步，交换 5 和 4，得到 123564；然后需要将 5 之后的数重置为升序，得到 123546。显然 123546 比 123564 更小，123546 就是 123465 的下一个排列
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.3 MB,击败了79.02% 的Go用户
func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	// 找到第一个降序的
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		// 找到[j, end]中从右往左找的第一个大于i的值
		for nums[i] >= nums[k] {
			k--
		}
		// 交换
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 反转[j, end]
	for a, b := j, len(nums)-1; a < b; a, b = a+1, b-1 {
		nums[a], nums[b] = nums[b], nums[a]
	}
}

// Search 33. 搜索旋转排序数组
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.4 MB,击败了76.37% 的Go用户
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// SearchRange 34. 在排序数组中查找元素的第一个和最后一个位置
// 二分法
//
//	执行耗时:3 ms,击败了92.84% 的Go用户
//	内存消耗:4.3 MB,击败了32.43% 的Go用户
func SearchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			// 找到了目标
			l := mid - 1
			r := mid + 1
			for ; l >= 0 && nums[l] == target; l-- {
			}
			for ; r < len(nums) && nums[r] == target; r++ {
			}
			l += 1
			r -= 1
			return []int{l, r}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}

// SearchInsert 35. 搜索插入位置
//
//	执行耗时:3 ms,击败了60.59% 的Go用户
//	内存消耗:2.8 MB,击败了75.85% 的Go用户
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// IsValidSudoku 36. 有效的数独
// 这个方法太吃计算量
//
//	执行耗时:8 ms,击败了5.35% 的Go用户
//	内存消耗:3.3 MB,击败了18.17% 的Go用户
func IsValidSudoku(board [][]byte) bool {
	otherMap := map[string]map[byte]bool{}
	for i := 0; i < len(board); i++ {
		// 这里是一列
		rowMap := map[byte]bool{}
		rolMap := map[byte]bool{}
		for j := 0; j < len(board); j++ {
			other := fmt.Sprintf("%d%d", i/3, j/3)
			curMap, ok := otherMap[other]
			if !ok {
				curMap = map[byte]bool{}
			}
			// 这里是斜线
			if board[i][j] != '.' {
				if _, ok := curMap[board[i][j]]; ok {
					return false
				}
				// 这里是一行
				if _, ok := rowMap[board[i][j]]; ok {
					return false
				}
				rowMap[board[i][j]] = true
				curMap[board[i][j]] = true
				otherMap[other] = curMap
			}
			if board[j][i] != '.' {
				// 这里是一列
				if _, ok := rolMap[board[j][i]]; ok {
					return false
				}
				rolMap[board[j][i]] = true
			}
		}
	}
	return true
}

// IsValidSudoku1 方法二
//
//	执行耗时:3 ms,击败了44.65% 的Go用户
//	内存消耗:2.5 MB,击败了99.58% 的Go用户
func IsValidSudoku1(board [][]byte) bool {
	var rows, cols [9][9]int
	var other [3][3][9]int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == '.' {
				continue
			}
			index := board[i][j] - '1'
			rows[i][index] += 1
			cols[j][index] += 1
			other[i/3][j/3][index] += 1
			if rows[i][index] > 1 || cols[j][index] > 1 || other[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

// CountAndSay 38. 外观数列
//
//	执行耗时:12 ms,击败了13.50% 的Go用户
//	内存消耗:7.3 MB,击败了17.50% 的Go用户
func CountAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		newS := ""
		first := s[0]
		num := 1
		for j := 1; j < len(s); j++ {
			if first == s[j] {
				num++
			} else {
				newS += fmt.Sprintf("%d%d", num, first-'0')
				num = 1
				first = s[j]
			}
		}
		newS += fmt.Sprintf("%d%d", num, first-'0')
		s = newS
	}
	return s
}

// CombinationSum 39. 组合总和
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
// 难点如何去重: 下一个循环不考虑上一个循环的数字
//
//	执行耗时:4 ms,击败了33.42% 的Go用户
//	内存消耗:4.1 MB,击败了11.56% 的Go用户
func CombinationSum(candidates []int, target int) [][]int {
	var r [][]int
	var fn func([]int, int, int)
	fn = func(arr []int, begin int, target int) {
		if target == 0 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			r = append(r, tmp)
			return
		}
		if target < 0 {
			return
		}
		for i := begin; i < len(candidates); i++ {
			fn(append(arr, candidates[i]), i, target-candidates[i])
		}
	}
	fn([]int{}, 0, target)
	return r
}

// CombinationSum2 组合总和2
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:4.6 MB,击败了7.85% 的Go用户
func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var r [][]int
	var fn func([]int, int, int)
	fn = func(arr []int, begin int, target int) {
		if target == 0 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			r = append(r, tmp)
			return
		}
		if target < 0 {
			return
		}
		for i := begin; i < len(candidates); i++ {
			if i > begin && candidates[i] == candidates[i-1] {
				continue
			}
			fn(append(arr, candidates[i]), i+1, target-candidates[i])
		}
	}
	fn([]int{}, 0, target)
	return r
}
