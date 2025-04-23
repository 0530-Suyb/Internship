package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr := []int{3, 2, 1, 9, 4, 5, 8, 10, 6, 7}
	// // quickSort(arr)
	// arr = mergeSort(arr)
	// fmt.Println(arr, inverNum)

	n, m := 0, 0
	fmt.Scanln(&n, &m)
	arr := make([][2]int, n*m)
	for i := 0; i < n*m; i++ {
		v := 0
		fmt.Scan(&v)
		arr[i] = [2]int{v, i}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	sum := 0
	cao := m * (m - 1) / 2
	for i := 0; i < n; i++ {
		count, _ := countInversionNum(arr[i*m : (i+1)*m])
		sum += cao - count // cao为组合数，count为逆序数，cao-count为正序数，只有正序时才会经过前面较小的数
	}

	fmt.Println(sum)
}

// 求逆序数
func countInversionNum(arr [][2]int) (invN int, arrNew [][2]int) {
	n := len(arr)
	if n <= 1 {
		return 0, arr
	}
	mid := n / 2
	left := arr[:mid]
	right := arr[mid:]
	c1, left := countInversionNum(left)
	c2, right := countInversionNum(right)
	fmt.Println(left, right)
	c3, arrNew := mergeAndCountInverNum(left, right)
	return c1 + c2 + c3, arrNew
}

func mergeAndCountInverNum(left, right [][2]int) (invN int, arrNew [][2]int) {
	arrTmp := make([][2]int, len(left)+len(right))
	i, j, k, count := 0, 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i][1] <= right[j][1] {
			arrTmp[k] = left[i]
			i++
		} else {
			arrTmp[k] = right[j]
			j++
			count += len(left) - i
		}
		k++
	}
	for i < len(left) {
		arrTmp[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arrTmp[k] = right[j]
		j++
		k++
	}
	return count, arrTmp
}

// 快速排序
// 时间复杂度：最优O(nlogn)，最差O(n^2)
// 空间复杂度：O(logn)
// 稳定性：不稳定
func quickSort(arr []int) {
	partition(arr, 0, len(arr)-1)
}

func partition(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	privot := arr[mid]
	for left < right {
		for arr[left] < privot {
			left++
		}
		for arr[right] > privot {
			right--
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	partition(arr, 0, left-1)
	partition(arr, left+1, right)
}

// 归并排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
// 稳定性：稳定
var inverNum int

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := mergeSort(arr[:len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	new := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			new = append(new, left[i])
			i++
		} else {
			new = append(new, right[j])
			j++
			inverNum += len(left) - i
		}
	}
	for i < len(left) {
		new = append(new, left[i])
		i++
	}
	for j < len(right) {
		new = append(new, right[j])
		j++
	}
	return new
}

// 二分查找
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
