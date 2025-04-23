package main

import (
	"fmt"
)

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	lA, rA := arr[:mid], arr[mid:]
	lA = MergeSort(lA)
	rA = MergeSort(rA)
	return merge(lA, rA)
}

func merge(a1, a2 []int) []int {
	newA := []int{}
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			newA = append(newA, a1[i])
			i++
		} else {
			newA = append(newA, a2[j])
			j++
		}
	}
	if i != len(a1) {
		newA = append(newA, a1[i:]...)
	} else {
		newA = append(newA, a2[j:]...)
	}
	return newA
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	l, r, pivot := 0, len(arr)-1, arr[0]
	for l < r {
		for l < r && arr[l] <= pivot {
			l++
		}
		for l < r && arr[r] > pivot {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	if arr[l] < arr[0] { // 注意对比交换！！！
		arr[0], arr[l] = arr[l], arr[0]
	}
	QuickSort(arr[:l])
	QuickSort(arr[l:])
}

func main() {
	fmt.Println(MergeSort([]int{4, 3, 2, 1, 5, 3, 9}))
	quickT := []int{4, 3, 2, 1, 5, 9, 0, 6, 3}
	QuickSort(quickT)
	fmt.Println(quickT)
}
