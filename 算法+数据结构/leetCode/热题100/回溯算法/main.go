package main

import (
	"fmt"
	"strconv"
)

/*
回溯模板

void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}
*/

// 1. 重组IP地址

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}
	// result := make([]string, 0, 10)
	// ip := make([]string, 0, 4)
	// backtracking(ip, s, &result)

	result := []string{}
	m := make(map[string]struct{})
	ip := make([]string, 0, 4)
	backtracking2(ip, s, m)
	for s, _ := range m {
		result = append(result, s)
	}

	return result
}

// 切记切片作为参数是值传递，函数内对切片的长度和容量的修改不影响到函数外，最终只有底层的数组会被改到
// 也可以考虑传map
func backtracking(ip []string, str string, result *[]string) {
	if len(ip) == 4 || len(str) == 0 {
		if len(str) == 0 && len(ip) == 4 {
			s := ip[0] + "." + ip[1] + "." + ip[2] + "." + ip[3]
			*result = append(*result, s)
		}
		return
	}

	for i := 1; i <= 3; i++ {
		if len(str) < i {
			break
		}
		n, _ := strconv.Atoi(str[:i])
		if n < 256 {
			ip = append(ip, str[:i])
			backtracking(ip, str[i:], result)
			if n == 0 { // 为0则不用在试两位数、三位数的情况，因为不能为前导零
				break
			} else {
				ip = ip[:len(ip)-1]
			}
		}
	}
}

func backtracking2(ip []string, str string, result map[string]struct{}) {
	if len(ip) == 4 || len(str) == 0 {
		if len(str) == 0 && len(ip) == 4 {
			s := ip[0] + "." + ip[1] + "." + ip[2] + "." + ip[3]
			result[s] = struct{}{}
		}
		return
	}

	for i := 1; i <= 3; i++ {
		if len(str) < i {
			break
		}
		n, _ := strconv.Atoi(str[:i])
		if n < 256 {
			ip = append(ip, str[:i])
			backtracking2(ip, str[i:], result)
			if n == 0 { // 为0则不用在试两位数、三位数的情况，因为不能为前导零
				break
			} else {
				ip = ip[:len(ip)-1]
			}
		}
	}
}

// 2. 全排列：不重复数组

func permute(nums []int) [][]int {
	arr := make([]int, len(nums))
	res := [][]int{}
	permute_dfs(nums, arr, 0, &res)
	return res
}

func permute_dfs(nums []int, arr []int, arrLen int, res *[][]int) { // 使用二维切片的引用，因为如果传对象只会进行拷贝，函数内二维切片的元素长度变化外部看不见
	if len(nums) == 0 {
		*res = append(*res, append([]int(nil), arr...)) // 注意是append([]int(nil), arr...)而非arr，因为如果将arr切片添加到res中，每次进入函数时arr的底层数组都是一样的，最终res中每个切片的底层数组都是同一个
		return
	}
	for i := 0; i < len(nums); i++ {
		arr[arrLen] = nums[i]
		nums[0], nums[i] = nums[i], nums[0]
		permute_dfs(nums[1:], arr, arrLen+1, res)
		nums[i], nums[0] = nums[0], nums[i]
	}
}

// 3. 全排列：重复数组

func permuteUnique(nums []int) [][]int {
	l := len(nums)
	res := [][]int{}
	flag := make([]bool, l)
	arr := make([]int, l)
	var dfs func(i int)
	dfs = func(i int) {
		if l == i {
			res = append(res, append([]int(nil), arr...))
			return
		}
		m := make(map[int]struct{})
		for idx, used := range flag {
			if !used {
				if _, ok := m[nums[idx]]; ok {
					continue
				}
				m[nums[idx]] = struct{}{}
				flag[idx] = true
				arr[i] = nums[idx]
				dfs(i + 1)
				flag[idx] = false
			}
		}
	}
	dfs(0)
	return res
}

func main() {
	fmt.Println("复原IP地址：25525511135")
	fmt.Println(restoreIpAddresses("25525511135")) // 复原IP地址
	fmt.Println("全排列：不重复数组")
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println("全排列：重复数组")
	fmt.Println(permuteUnique([]int{1, 1, 3}))
}
