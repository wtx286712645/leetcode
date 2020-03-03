package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	l1 := buildList([]int{1,2,3,4,5})
	l2 := buildList([]int{1,2,9,4,5})
	printList(l1)
	printList(l2)

	ret := AddTwoNumbers(l1, l2)
	printList(ret)
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	list := &ListNode{
		Val:nums[0],
		Next:nil,
	}

	tail := list
	if len(nums) > 1 {
		for i:=1; i<len(nums); i++ {
			tail.Next = &ListNode{
				Val:nums[i],
				Next:nil,
			}
			tail = tail.Next
		}
	}

	return list
}

func printList(list *ListNode)  {
	for list != nil {
		fmt.Printf("%v-", list.Val)
		list = list.Next
	}
	fmt.Printf("\n")
}
