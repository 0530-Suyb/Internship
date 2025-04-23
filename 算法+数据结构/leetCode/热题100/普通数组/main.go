package main

import (
	"fmt"
)

// 1. 最大子数组和

// 动态规划
// f(i)表示以i索引结尾的最大子数组和，则状态转移方程f(i)=max{f(i-1)+nums[i], nums[i]}
// 初始f(0)=nums[0]
func maxSubArray(nums []int) int {
	pre, res := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		pre = max(pre+nums[i], nums[i])
		res = max(res, pre)
	}
	return res
}

// 分治法
func maxSubArray2(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(lS, rS Status) (s Status) {
	s.iSum = lS.iSum + rS.iSum
	s.lSum = max(lS.lSum, lS.iSum+rS.lSum)
	s.rSum = max(rS.rSum, rS.iSum+lS.rSum)
	s.mSum = max(max(lS.mSum, rS.mSum), lS.rSum+rS.lSum)
	return
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lStatus := get(nums, l, m)
	rStatus := get(nums, m+1, r)
	return pushUp(lStatus, rStatus)
}

type Status struct {
	iSum, lSum, rSum, mSum int
}

func main() {
	fmt.Println(`1. 最大子数组和: [-2,1,-3,4,-1,2,1,-5,4]`)
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
