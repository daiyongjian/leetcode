package main

import (
	"code/code/a111"
	"fmt"
)

func main() {
	//input := 'AB'
	//input := '  +  413'
	//input := 1994
	//input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//input := '{[]}'
	//input := 2
	//input := [][]byte{
	//	{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
	//	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	//	{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
	//	{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
	//	{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
	//	{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
	//	{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
	//	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	//	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	//}
	//input := [][]int{
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9, 10, 11, 12},
	//}
	//input := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	//input := []int{2, 3, 1, 1, 4}
	//input := [][]int{
	//	{1, 3},
	//	{2, 6},
	//	{8, 10},
	//	{15, 18},
	//}
	//input := 4
	//input := &a81.ListNode{Val: 1, Next: &a81.ListNode{Val: 2, Next: &a81.ListNode{Val: 3, Next: &a81.ListNode{Val: 4, Next: &a81.ListNode{Val: 5, Next: &a81.ListNode{Val: 6}}}}}}
	//input := &a61.ListNode{Val: 1, Next: &a61.ListNode{Val: 2}}
	//input := [][]byte{
	//	{'A', 'B', 'C', 'E'},
	//	{'S', 'F', 'C', 'S'},
	//	{'A', 'D', 'E', 'E'},
	//}
	//input := []int{1, 2, 2}
	//input := "25525511135"
	//input := &a101.TreeNode{Val: 3, Left: &a101.TreeNode{Val: 9}, Right: &a101.TreeNode{Val: 20, Left: &a101.TreeNode{Val: 15}, Right: &a101.TreeNode{Val: 7}}}
	input := 4
	output := a111.Generate(input)
	//fmt.Printf('输入: %s', input)
	//fmt.Printf('输出: %s', output)
	fmt.Printf("输入: %v\n", input)
	//for output != nil {
	//	fmt.Printf("%d,", output.Val)
	//	output = output.Next
	//}
	fmt.Printf("输出: %v\n", output)
}
