package a81

import "sort"

// Search 搜索旋转排序数组II
// [1,0,1,1,1] 0
//
//	执行耗时:4 ms,击败了55.74% 的Go用户
//	内存消耗:3 MB,击败了100.00% 的Go用户
func Search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[left] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[mid] >= nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除排序链表中的重复元素II
// [1,2,3,3,4,4,5]
//
//	执行耗时:3 ms,击败了60.41% 的Go用户
//	内存消耗:2.7 MB,击败了99.61% 的Go用户
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fake := &ListNode{
		Next: head,
	}
	cur := fake
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return fake.Next
}

// DeleteDuplicates1 删除排序链表中的重复元素
// 需要注意最后可能会导致cur变成nil
//
//	执行耗时:3 ms,击败了60.49% 的Go用户
//	内存消耗:2.8 MB,击败了72.19% 的Go用户
func DeleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}

// Partition 分隔链表
// [1,4,3,2,5,2]
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.2 MB,击败了75.77% 的Go用户
func Partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	leftH := &ListNode{}
	rightH := &ListNode{}
	cur := head
	left := leftH
	right := rightH
	for cur != nil {
		if cur.Val >= x {
			right.Next = cur
			right = right.Next
		} else {
			left.Next = cur
			left = left.Next
		}
		tmp := cur.Next
		cur.Next = nil
		cur = tmp
	}
	left.Next = rightH.Next
	return leftH.Next
}

// Merge 合并两个有序数组
// nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.2 MB,击败了46.50% 的Go用户
func Merge(nums1 []int, m int, nums2 []int, n int) {
	left, right := m-1, n-1
	for i := len(nums1) - 1; i >= 0; i-- {
		if right < 0 {
			nums1[i] = nums1[left]
			left--
		} else if left < 0 {
			nums1[i] = nums2[right]
			right--
		} else {
			if nums1[left] > nums2[right] {
				nums1[i] = nums1[left]
				left--
			} else {
				nums1[i] = nums2[right]
				right--
			}
		}
	}
}

// GrayCode 格雷编码
func GrayCode(n int) []int {
	ans := make([]int, 1, 1<<n)
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]|1<<(i-1))
		}
	}
	return ans
}

// SubsetsWithDup 子集II
//
//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.3 MB,击败了80.79% 的Go用户
func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var r [][]int
	var ans []int
	var fn func(int)
	vis := make([]bool, len(nums))
	fn = func(cur int) {
		tmp := make([]int, len(ans))
		copy(tmp, ans)
		r = append(r, tmp)
		if cur == len(nums) {
			return
		}
		for i := cur; i < len(nums); i++ {
			if vis[i] || i > 0 && !vis[i-1] && nums[i] == nums[i-1] {
				continue
			}
			vis[i] = true
			ans = append(ans, nums[i])
			fn(i + 1)
			vis[i] = false
			ans = ans[:len(ans)-1]
		}
	}
	fn(0)
	return r
}
