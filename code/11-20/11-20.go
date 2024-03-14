package _1_20

import (
	"sort"
)

// MaxArea 11: 盛最多水的容器
//
//	 双指针
//		执行耗时:64 ms,击败了62.99% 的Go用户
//		内存消耗:8.2 MB,击败了17.95% 的Go用户
func MaxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		var curArea int
		if height[left] > height[right] {
			curArea = height[right] * (right - left)
			right--
		} else {
			curArea = height[left] * (right - left)
			left++
		}
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}

// IntToRoman 12: 整数转罗马数字
// hash表
// 字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
// 1 <= num <= 3999
//
//	执行耗时:12 ms,击败了22.19% 的Go用户
//	内存消耗:3 MB,击败了57.94% 的Go用户
func IntToRoman(num int) string {
	r := ""
	// 先算千
	num1 := num / 1000
	if num1 != 0 {
		for i := 0; i < num1; i++ {
			r += "M"
		}
	}
	// 再算百
	r += intToRomanForUnit((num%1000)/100, "C", "D", "M")
	r += intToRomanForUnit((num%100)/10, "X", "L", "C")
	r += intToRomanForUnit(num%10, "I", "V", "X")
	return r
}

func intToRomanForUnit(unit int, num1, num2, num3 string) string {
	if unit == 0 {
		return ""
	}
	r := ""
	if unit < 4 {
		for i := 0; i < unit; i++ {
			r += num1
		}
	} else if unit == 4 {
		r += num1 + num2
	} else if unit == 5 {
		r += num2
	} else if unit == 9 {
		r += num1 + num3
	} else {
		r += num2
		unit -= 5
		for i := 0; i < unit; i++ {
			r += num1
		}
	}
	return r
}

// LongestCommonPrefix 14. 最长公共前缀
// 没啥算法，强写
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.2 MB,击败了78.89% 的Go用户
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	firstWord := strs[0]
	for i := 0; i < len(firstWord); i++ {
		// firstLetter
		firstLetter := strs[0][i]
		isPrefix := true
		for _, str := range strs[1:] {
			if i >= len(str) {
				isPrefix = false
				break
			}
			if firstLetter != str[i] {
				isPrefix = false
				break
			}
		}
		if !isPrefix {
			return strs[0][:i]
		}
	}
	return firstWord
}

// ThreeSum 15: 三数之和
// 难点1： 解决重复的问题
// 所以需要先排序
//
// 难点2: 三循环太久
// 利用排序优势，不需要完全循环
//
//	执行耗时:44 ms,击败了62.92% 的Go用户
//	内存消耗:8 MB,击败了83.02% 的Go用户
func ThreeSum(nums []int) [][]int {
	curLen := len(nums)
	sort.Ints(nums)
	var r [][]int
	for i := 0; i < curLen-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := curLen - 1
		target := 0 - nums[i]
		for j := i + 1; j < curLen-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			if k == j {
				break
			}
			if nums[j]+nums[k] == target {
				r = append(r, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return r
}

// ThreeSumClosest 16: 最接近的三数之和
// eg: [-1,2,1,-4] 1 结果2
// 双指针
//
//	执行耗时:14 ms,击败了35.81% 的Go用户
//	内存消耗:2.7 MB,击败了92.98% 的Go用户
func ThreeSumClosest(nums []int, target int) int {
	curLen := len(nums)
	sort.Ints(nums)
	minNum := nums[0] + nums[1] + nums[2]
	for i := 0; i < curLen-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := curLen - 1
		j := i + 1
		for j < k {
			curSum := nums[i] + nums[j] + nums[k]
			if curSum == target {
				return curSum
			}
			if abs(curSum-target) < abs(minNum-target) {
				minNum = curSum
			}
			// 如果太大了，就往小移动
			if curSum > target {
				k0 := k - 1
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				// 如果小了，就往大移动
				j0 := j + 1
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return minNum
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// LetterCombinations 17: 电话号码的字母组合
// 剪枝算法
//
//	执行耗时:1 ms,击败了13.19% 的Go用户
//	内存消耗:1.9 MB,击败了23.21% 的Go用户
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	var r []string
	var fn func(int, string)
	fn = func(i int, s string) {
		if i == len(digits) {
			r = append(r, s)
			return
		}
		target := m[digits[i]]
		for _, item := range target {
			fn(i+1, s+item)
		}
	}
	fn(0, "")
	return r
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// FourSum 18: 四数之和
// 双指针 + 剪枝
//
//	执行耗时:4 ms,击败了89.11% 的Go用户
//	内存消耗:2.6 MB,击败了94.17% 的Go用户
func FourSum(nums []int, target int) [][]int {
	var r [][]int
	sort.Ints(nums)
	curLen := len(nums)
	for i := 0; i < curLen-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[curLen-1]+nums[curLen-2]+nums[curLen-3] < target {
			continue
		}
		for j := i + 1; j < curLen-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[curLen-1]+nums[curLen-2] < target {
				continue
			}
			for left, right := j+1, curLen-1; left < right; {
				curSum := nums[i] + nums[j] + nums[left] + nums[right]
				if curSum == target {
					r = append(r, []int{nums[i], nums[j], nums[left], nums[right]})
					// left和right都找到下一个值
					for left++; left < right && nums[left] == nums[left-1]; left++ {

					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {

					}
				} else if curSum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return r
}

// RemoveNthFromEnd 19: 删除链表的倒数第N个节点
// 双指针
// 注意1: 需要一个fake指针，因为head可能也会被删除
// 注意2: n2本身可能为nil
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.1 MB,击败了76.83% 的Go用户
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fake := &ListNode{Next: head}
	n1 := fake
	n2 := fake
	if head == nil || head.Next == nil {
		return nil
	}
	for i := 0; i < n; i++ {
		n2 = n2.Next
	}
	for n2 != nil && n2.Next != nil {
		n1 = n1.Next
		n2 = n2.Next
	}
	// 此时的n1就是倒数第N个节点
	n1.Next = n1.Next.Next
	return fake.Next
}

// IsValid 20. 有效的括号
// 需要一个堆栈，先进后出
//
//	执行耗时:1 ms,击败了14.13% 的Go用户
//	内存消耗:1.9 MB,击败了76.66% 的Go用户
func IsValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			// 从stack中取一个
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if cur == '(' && s[i] == ')' {
				continue
			} else if cur == '[' && s[i] == ']' {
				continue
			} else if cur == '{' && s[i] == '}' {
				continue
			}
			return false
		}
	}
	return len(stack) == 0
}
