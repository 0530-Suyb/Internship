package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i, j := l, r
	for i < j {
		for nums[i] < partition {
			i++
		}
		for nums[j] > partition {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if nums[j] > partition {
		j--
	}
	if k > j {
		return quickSelect(nums, j+1, r, k)
	} else {
		return quickSelect(nums, l, j, k)
	}
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
