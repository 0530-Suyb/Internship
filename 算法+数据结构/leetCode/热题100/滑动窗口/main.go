package main

import (
	"fmt"
)

// 《无重复字符的最长子串》
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
// 示例 1:
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//
//	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// 提示：
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
//
// 时间复杂度：O(n)，快慢指针各遍历字符串一次
// 空间复杂度：O(n)
func lengthOfLongestSubstring(s string) int {
	slow, fast, res := 0, 0, 0
	m := make(map[byte]int)
	for fast < len(s) {
		if v, ok := m[s[fast]]; ok {
			for slow <= v {
				delete(m, s[slow])
				slow++
			}
		}
		m[s[fast]] = fast
		res = max(res, fast-slow+1)
		fast++
	}
	return res
}

// 《找到字符串中所有字母异位词》
// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 示例 1:
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//
// 示例 2:
// 输入: s = "abab", p = "ab"
// 输出: [0,1,2]
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
//
// 提示:
// 1 <= s.length, p.length <= 3 * 104
// s 和 p 仅包含小写字母

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	arr := []int{}
	isAna := func(subs string, p string) bool {
		a := [26]int{}
		for i := 0; i < len(subs); i++ {
			a[subs[i]-'a']++
			a[p[i]-'a']--
		}
		for i := 0; i < len(subs); i++ {
			if a[subs[i]-'a'] != 0 {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(s)-len(p)+1; i++ {
		if isAna(s[i:i+len(p)], p) {
			arr = append(arr, i)
		}
	}
	return arr
}

func findAnagrams2(s string, p string) (res []int) {
	if len(s) < len(p) {
		return res
	}

	a1, a2 := [26]int{}, [26]int{}
	for _, c := range p {
		a1[c-'a']++
	}

	for i, c := range s {
		a2[c-'a']++
		if k := 0; i >= len(p)-1 {
			for ; k < 26 && a1[k] == a2[k]; k++ {
			}
			if k == 26 {
				res = append(res, i-len(p)+1)
			}
			a2[s[i-len(p)+1]-'a']--
		}
	}

	return res

	// if len(s) < len(p) {
	// 	return []int{}
	// }

	// a1 := [26]int{}
	// a2 := [26]int{}
	// arr := []int{}
	// for i := 0; i < len(p); i++ {
	// 	a1[p[i]-'a']++
	// }

	// for i, j := 0, 0; i < len(s); i++ {
	// 	a2[s[i]-'a']++
	// 	if i+1 >= len(p) {
	// 		for k:=0; k<len(a2); k++{
	// 			if a1[k] != a2[k] {
	// 				break
	// 			}
	// 			if k == len(a2) - 1 {
	// 				arr = append(arr, j)
	// 			}
	// 		}
	// 		a2[s[j]-'a']--
	// 		j++
	// 	}
	// }
	// return arr
}

func main() {
	fmt.Println(`1. 无重复字符的最长子串长度："abcabcbb"`)
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(`2. 找到字符串中所有字母异位词: s = "cbaebabacd", p = "abc"`)
	fmt.Println(findAnagrams2("cbaebabacd", "abc"))
}
