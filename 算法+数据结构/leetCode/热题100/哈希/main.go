package main

import (
	"fmt"
)

// 1. 字母异位词分组

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	res := [][]string{}
	for _, s := range strs {
		k := insertSort(s)
		m[k] = append(m[k], s)
	}
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func insertSort(str string) string {
	newStr := make([]byte, len(str))
	copy(newStr, str)
	for i := 1; i < len(str); i++ {
		for j := i - 1; j >= 0; j-- {
			if newStr[j+1] < newStr[j] {
				newStr[j+1], newStr[j] = newStr[j], newStr[j+1]
			}
		}
	}
	return string(newStr)
}

// 2. 最长连续子序列

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	maxLen := 0
	for _, num := range nums {
		m[num] = true
	}
	for k := range m {
		i := 1
		for ; !m[k-1] && m[k+i]; i++ {
		}
		if i > maxLen {
			maxLen = i
		}
	}
	return maxLen
}

func main() {
	fmt.Println(`1. 字母异位词分组：["abc", "bca", "ete", "eet"]`)
	fmt.Println(groupAnagrams([]string{"abc", "bca", "ete", "eet"}))
	fmt.Println(`2. 最长连续子序列：[100,4,200,1,3,2]`)
	fmt.Println(longestConsecutive([]int{100, 2, 300, 3, 4, 1}))
}
