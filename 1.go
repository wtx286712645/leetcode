package leetcode

func TwoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i := 0 ; i < len(nums); i++ {
		container := target - nums[i]
		if _, ok := m[container]; ok {
			return []int{m[container], i}
		}
		m[nums[i]] = i
	}
	return nil
}

