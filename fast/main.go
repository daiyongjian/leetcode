package main

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

func main() {
	aa := []int{1}
	searchRange(aa, 1)
}
