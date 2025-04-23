package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 最大堆需要改比较为>
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 注意：一定是末尾元素是堆顶
func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}


// 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	hSize := len(nums)
	buildMaxHeap(nums, hSize)
	for i:=hSize-1; i>=hSize-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, 0, i)
	}
	return nums[0]
}

func buildMaxHeap(a []int, hSize int)  {
	for i:=(hSize)/2-1; i>=0; i-- {
		maxHeapify(a, i, hSize)
	}
}

func maxHeapify(a []int, i, hSize int) {
	l, r, largest := 2*i+1, 2*i+2, i
	if l<hSize && a[l]>a[largest] {
		largest = l
	}
	if r<hSize && a[r]>a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, hSize)
	}
}

/*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
 
提示：
1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
 
进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
*/

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	arr := []int{}
	for k := range m {
		arr = append(arr, k)
	}
	hSize := len(arr)
	var maxHeapify func(i, hSize int)
	maxHeapify = func(i, hSize int) {
		l, r, largest := 2*i+1, 2*i+2, i
		if l<hSize && m[arr[l]] > m[arr[largest]] {
			largest = l
		}
		if r<hSize && m[arr[r]] > m[arr[largest]] {
			largest = r
		}
		if largest != i {
			arr[largest], arr[i] = arr[i], arr[largest]
			maxHeapify(largest, hSize)
		}
	}
	for i:=hSize/2-1; i>=0; i-- {
		maxHeapify(i, hSize)
	}

	res := make([]int, k)
	for i:=0; i<k; i++ {
		res[i] = arr[0]
		arr[0], arr[hSize-1] = arr[hSize-1], arr[0]
		hSize--
		maxHeapify(0, hSize)
	}
	return res
}

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 6)
	// fmt.Println(heap.Pop(h))
	// fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3}, 2))
}

