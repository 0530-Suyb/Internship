package main

import (
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {
	var arr [26]int
	var a rune = 'a'
	for _, v := range s {
		arr[v-a]++
	}
	for _, v := range t {
		arr[v-a]--
	}
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int)
	for _, v := range s {
		sMap[v]++
	}
	for _, v := range t {
		sMap[v]--
		if sMap[v] < 0 {
			return false
		}
	}
	return true
}

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}) // struct{}不占用内存空间，set的值只有有或无，而如果用bool，会占用内存空间，结果为true、false、无
	for _, v := range nums1 {
		set[v] = struct{}{}
	}
	var res []int
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}

func isHappy(n int) bool {
	getSum := func(n int) (sum int) {
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		return
	}

	m := make(map[int]struct{})
	for n != 1 {
		m[n] = struct{}{}
		n = getSum(n)
		if _, ok := m[n]; ok {
			break
		}
	}
	return n == 1
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	m := make(map[int]int) // key不存在则返回零值，此处是0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += m[-v3-v4]
		}
	}
	return count
}

func canConstruct(ransomNote string, magazine string) bool {
	arr := [26]int{}
	for _, v := range magazine {
		arr[v-'a']++
	}
	for _, v := range ransomNote {
		arr[v-'a']--
		if arr[v-'a'] < 0 {
			return false
		}
	}
	return true
}

// error 超时
func threeSum(nums []int) [][]int {
	var res [][]int
	m := make(map[int]map[int]map[int]struct{})

	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)

	m1 := make(map[int][][2]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			m1[nums[i]+nums[j]] = append(m1[nums[i]+nums[j]], [2]int{i, j})
		}
	}
	for i := 0; i < len(nums); i++ {
		if v, ok := m1[-nums[i]]; ok {
			for _, arr := range v {
				if arr[0] == i || arr[1] == i {
					continue
				}
				arr1 := []int{nums[arr[0]], nums[arr[1]], nums[i]}
				sort.Ints(arr1)
				if m[arr1[0]] == nil {
					m[arr1[0]] = make(map[int]map[int]struct{})
					m[arr1[0]][arr1[1]] = make(map[int]struct{})
					m[arr1[0]][arr1[1]][arr1[2]] = struct{}{}
					res = append(res, arr1)
				} else if m[arr1[0]][arr1[1]] == nil {
					m[arr1[0]][arr1[1]] = make(map[int]struct{})
					m[arr1[0]][arr1[1]][arr1[2]] = struct{}{}
					res = append(res, arr1)
				} else if _, ok := m[arr1[0]][arr1[1]][arr1[2]]; !ok {
					m[arr1[0]][arr1[1]][arr1[2]] = struct{}{}
					res = append(res, arr1)
				}
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 去重
			continue
		}
		m := make(map[int]struct{})
		for j := i + 1; j < len(nums); j++ {
			if j > i+2 && nums[j] == nums[j-1] && nums[j-1] == nums[j-2] { // 去重，注意从j>i+2，且nums[j-1] == nums[j-2]
				continue
			}
			target := 0 - nums[i] - nums[j]
			if _, ok := m[target]; ok {
				res = append(res, []int{nums[i], nums[j], target})
				delete(m, target)
			} else {
				m[nums[j]] = struct{}{}
			}
		}
	}
	return res
}

func threeSum3(nums []int) [][]int {
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
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == n2 { // 注意是n2而非nums[l+1]
					l++
				}
				for l < r && nums[r] == n3 { // 注意是n3而非nums[r-1]
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

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if nums[i] > target && nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if nums[i]+nums[j] > target && nums[i]+nums[j] > 0 {
				continue
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				sum := nums[i] + nums[j] + n3 + n4
				if sum == target {
					res = append(res, []int{nums[i], nums[j], n3, n4})
					for l < r && nums[l] == n3 {
						l++
					}
					for l < r && nums[r] == n4 {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}

// 《最长连续序列》
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1：
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
//
// 示例 3：
// 输入：nums = [1,0,1,2]
// 输出：3
//
// 提示：
// 0 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	maxLen := 0
	for _, num := range nums {
		m[num] = true
	}
	for k := range m {
		i := 1
		if !m[k-1] { // ！！！去重
			for {
				if m[k+i] {
					i++
				} else {
					break
				}
			}
			if i > maxLen {
				maxLen = i
			}
		}

	}
	return maxLen
}

func main() {
	// a := 'a'
	// var b byte = 'b'
	// fmt.Printf("%T, %T", a, b)

	// fmt.Println(isAnagram2("anagram", "nagaram"))
	// fmt.Println(isAnagram2("好的呢", "的呢好"))
	// fmt.Println(isAnagram2("好的呢", "不好呢"))

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
