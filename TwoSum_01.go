package Leetcode_Go

//Given an array of integers, return indices of the two numbers such that they a
//dd up to a specific target.
//
// You may assume that each input would have exactly one solution, and you may n
//ot use the same element twice.
//
// Example:
//
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].
//
// Related Topics 数组 哈希表
// 👍 8695 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	res := []int{}
	for i, value1 := range nums {
		for j, value2 := range nums[i+1:] {
			if value1+value2 == target {
				res = append(res, i, j+i+1)
			}
		}

	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
