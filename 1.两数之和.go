/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	memo := make(map[int]int, 0)
	for k, v  := range nums {
		if index, ok := memo[target-v]; ok {
			return []int{index, k}
		}else{
			memo[v] = k
		}
	}
	return []int{}
}
// @lc code=end

