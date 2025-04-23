package main

import "fmt"

// 和为K的子数组
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
// 子数组是数组中元素的连续非空序列。
//
// 示例 1：
// 输入：nums = [1,1,1], k = 2
// 输出：2
//
// 示例 2：
// 输入：nums = [1,2,3], k = 3
// 输出：2
//
// 提示：
// 1 <= nums.length <= 2 * 104
// -1000 <= nums[i] <= 1000
// -107 <= k <= 107

// 前缀和
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func subarraySum(nums []int, k int) int {
	// 不是滑动窗口，是前缀和
	sum := 0
	if len(nums) != 0 && nums[0] == k {
		sum++
	}
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
		if nums[i] == k {
			sum++
		}
		for j := 0; j < i; j++ {
			if nums[i]-nums[j] == k {
				sum++
			}
		}
	}
	return sum
}

// 前缀和 + 哈希表优化
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func subarraySum2(nums []int, k int) int {
	res, pre := 0, 0
	m := map[int]int{0: 1}
	for _, n := range nums {
		pre += n
		if v, ok := m[pre-k]; ok {
			res += v
		}
		m[pre]++
	}
	return res
	// sum, pre := 0, 0
	// m := make(map[int]int)
	// m[0] = 1 // 考虑nums[n]==k的情况
	// for i := 0; i < len(nums); i++ {
	// 	pre += nums[i]
	// 	if _, ok := m[pre-k]; ok {
	// 		sum += m[pre-k]
	// 	}
	// 	m[pre]++
	// }
	// return sum
}

func main() {
	fmt.Println(`1. 和为K的子数组: nums = [1,2,3], k = 3`)
	fmt.Println(subarraySum2([]int{1,2,3}, 3))
}
