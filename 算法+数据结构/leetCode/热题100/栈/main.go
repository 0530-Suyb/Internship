package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString2(s string) string {
	str := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			start := ptr
			for s[ptr] >= '0' && s[ptr] <= '9' {
				ptr++
			}
			str = append(str, string(s[start:ptr]))
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			str = append(str, string(cur))
			ptr++
		} else {
			start := len(str) - 1
			for str[start] != "[" {
				start--
			}
			subStr := append([]string{}, str[start+1:len(str)]...)
			sub := strings.Join(subStr, "")
			repT, _ := strconv.Atoi(str[start-1])
			str = str[:start-1]
			str = append(str, strings.Repeat(sub, repT))
		}
	}
	return strings.Join(str, "")
}

func decodeString(s string) string {
	if len(s) < 2 {
		return s
	}
	res := ""
	ptr := 0
	for ptr < len(s) {
		if s[ptr] >= '0' && s[ptr] <= '9' {
			start := ptr
			for s[ptr] >= '0' && s[ptr] <= '9' {
				ptr++
			}
			digits := s[start:ptr]
			repT, _ := strconv.Atoi(digits)
			start = ptr + 1
			stack := []byte{}
			for {
				if s[ptr] == '[' {
					stack = append(stack, '[')
				} else if s[ptr] == ']' {
					stack = stack[:len(stack)-1]
				}
				ptr++
				if len(stack) == 0 {
					break
				}
			}
			subStr := s[start : ptr-1]
			fmt.Println("subStr: ", subStr)
			res += strings.Repeat(decodeString(subStr), repT)
		} else {
			res += string(s[ptr])
			ptr++
		}
	}
	return res
}

/*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
示例 2:

输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]


提示：

1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
*/

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}
	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		top := len(stack) - 1
		for top >= 0 && temperatures[stack[top]] < temperatures[i] {
			res[stack[top]] = i - stack[top]
			top--
		}
		stack = stack[:top+1]
		stack = append(stack, i)
	}
	return res
}

func main() {
	// fmt.Println(decodeString2("cc10[2[ab]dd]cc"))
	// fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
}
