package _1_30

import "math"

// 21: MergeTwoLists 合并两个有序的列表
//
//	执行耗时:3 ms,击败了26.85% 的Go用户
//	内存消耗:2.4 MB,击败了6.29% 的Go用户
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	h := &ListNode{}
	r := h
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			r.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			r.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		r = r.Next
	}
	if list1 != nil {
		r.Next = list1
	}
	if list2 != nil {
		r.Next = list2
	}
	return h.Next
}

// GenerateParenthesis 括号生成
// 每个字符串两种结果，"("和")"
// 回溯 + 剪枝
//
//	执行耗时:2 ms,击败了28.66% 的Go用户
//	内存消耗:2.6 MB,击败了63.35% 的Go用户
func GenerateParenthesis(n int) []string {
	var r []string
	var fn func(string, int, int)
	fn = func(s string, left int, right int) {
		if len(s) == 2*n {
			r = append(r, s)
			return
		}
		// 无非两种选择
		if left < n {
			fn(s+"(", left+1, right)
		}
		if right < left {
			fn(s+")", left, right+1)
		}
	}
	fn("", 0, 0)
	return r
}

// SwapPairs 两两交换链表中的节点
// 需要保留一个临时指针永远指向分段的开头
//
//	执行耗时:1 ms,击败了12.63% 的Go用户
//	内存消耗:2 MB,击败了76.12% 的Go用户
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	r := head.Next
	pre := &ListNode{Next: head}
	for pre.Next != nil && pre.Next.Next != nil {
		start := pre.Next
		end := pre.Next.Next
		pre.Next = end
		start.Next = end.Next
		end.Next = start
		pre = start
	}
	return r
}

// RemoveDuplicates 26: 删除有序数组中的重复项
//
//	执行耗时:2 ms,击败了94.56% 的Go用户
//	内存消耗:4.3 MB,击败了99.60% 的Go用户
func RemoveDuplicates(nums []int) int {
	curLen := len(nums)
	left := 1
	for i := 1; i < curLen; i++ {
		if nums[i] != nums[i-1] {
			nums[left] = nums[i]
			// 这里已经加了1了，所以结果就不用加1了
			left++
		}
	}
	return left
}

// RemoveElement 27. 移除元素
//
//	执行耗时:1 ms,击败了19.74% 的Go用户
//	内存消耗:2 MB,击败了31.22% 的Go用户
func RemoveElement(nums []int, val int) int {
	curLen := len(nums)
	left := 0
	for i := 0; i < curLen; i++ {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

// RemoveElement1 27. 移除元素
// 双指针 优化了点内存
//
//	执行耗时:1 ms,击败了19.74% 的Go用户
//	内存消耗:2 MB,击败了75.75% 的Go用户
func RemoveElement1(nums []int, val int) int {
	curLen := len(nums)
	left := 0
	right := curLen
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}

// StrStr 28. 找出字符串第一个匹配项的下标
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:1.9 MB,击败了99.75% 的Go用户
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if needle[0] == haystack[i] {
			j := 1
			for ; j < len(needle) && needle[j] == haystack[i+j]; j++ {
			}
			if j == len(needle) {
				return i
			}
		}
	}
	return -1
}

// 快速乘
// x 和 y 是负数，z 是正数
// 判断 z * y >= x 是否成立
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 { // 不能使用除法
		if z&1 > 0 {
			// 需要保证 result + add >= x
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 {
			// 需要保证 add + add >= x
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

func divide(dividend, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1 // 注意溢出，并且不能使用除法
		if quickAdd(divisor, mid, dividend) {
			ans = mid
			if mid == math.MaxInt32 { // 注意溢出
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans
}
