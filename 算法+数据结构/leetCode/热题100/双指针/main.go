package main

import (
	"fmt"
	"sort"
)

// 1. 盛最多水的容器
func maxArea(height []int) int {
	maxA, fast := 0, 0
	for ; fast < len(height); fast++ {
		for i := 0; i < fast; i++ {
			area := (fast - i) * min(height[i], height[fast])
			if area > maxA {
				maxA = area
			}
			if height[i] > height[fast] {
				break
			} // 无需再比后续的了，因为往后如果左高于右边界则宽度在减少，左低于右边界则高度和宽度都在减少
		}
	}
	return maxA
}

func maxArea2(height []int) int {
	l, r, maxA := 0, len(height)-1, 0
	for l < r {
		area := (r - l) * min(height[l], height[r])
		maxA = max(area, maxA)
		if height[l] < height[r] {
			l++
		} else { // 优化：右边界为较低时往左移，且新边界若更低则无需再比较
			for tmp := height[r]; height[r] <= tmp && l < r; r-- {
			}
		}
	}
	return maxA
}

// 2. 三数之和

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			sum := nums[i] + n2 + n3
			if sum == 0 {
				res = append(res, []int{nums[i], n2, n3})
				for l < r && nums[l] == n2 { // 注意是n2而非nums[l]
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// 3. 接雨水

// 动态规划
func trap(height []int) int {
	lenH := len(height)
	if lenH < 3 {
		return 0
	}

	leftMax := make([]int, lenH)
	rightMax := make([]int, lenH)
	leftMax[0] = height[0]
	rightMax[lenH-1] = height[lenH-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
		rightMax[lenH-i-1] = max(rightMax[lenH-i], height[lenH-i-1])
	}

	res := 0
	for i := 0; i < lenH; i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}

	return res
}

// 单调栈
func trap2(height []int) int {
	stack := []int{}
	res := 0
	for i:=0; i<len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			res += (i-left-1) * (min(height[i], height[left]) - height[top])
		}
		stack = append(stack, i)
	}
	return res
}

// 双指针
func trap3(height []int) int {
	l, r, leftMax, rightMax := 0, len(height)-1, 0, 0
	res := 0
	for l<r {
		leftMax = max(leftMax, height[l])
		rightMax = max(rightMax, height[r])
		if height[l] < height[r] {
			res += leftMax - height[l]
			l++
		} else {
			res += rightMax - height[r]
			r--
		}
	}
	return res
}

func main() {
	fmt.Println(`1. 盛最多水的容器: [1, 8, 6, 2, 5, 4, 8, 3, 7]`)
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(`2. 三数之和: [-1, 0, 1, 2, -1, -4]`)
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(`3. 接雨水: [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]`)
	fmt.Println(trap3([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
