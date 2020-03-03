package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{1,2,3,4,5,6}
	target1 := 8

	ret1 := TwoSum(nums, target1)
	fmt.Println(ret1)
}
