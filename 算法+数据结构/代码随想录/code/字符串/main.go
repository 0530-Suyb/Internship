package main

import (
	"fmt"
	"strings"
)

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i] ^= s[j]
		s[j] ^= s[i]
		s[i] ^= s[j]
		i++
		j--
	}
}

func reverseStr(s string, k int) string {
	reverse := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; {
			s[i] ^= s[j]
			s[j] ^= s[i]
			s[i] ^= s[j]
			i++
			j--
		}
	}

	str := []byte(s)

	for i := 0; i < len(str); i += (2 * k) {
		if i+k < len(str) {
			reverse([]byte(str[i : i+k]))
		} else {
			reverse([]byte(str[i:]))
		}
	}

	return string(str)
}

func replaceNumber() {
	s := []byte{}
	fmt.Scanf("%s", &s)
	insertElement := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' {
			s = append(s[:i], append(insertElement, s[i+1:]...)...)
			i += len(insertElement) - 1
		}
	}
	fmt.Println(string(s))
}

func replaceNumber2() {
	s := []byte{}
	fmt.Scanln(&s)
	len1 := len(s)
	n := 0
	insertElement := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
	len2 := len(insertElement)
	for _, c := range s {
		if c <= '9' && c >= '0' {
			n++
		}
	}
	for i := 0; i < n; i++ {
		s = append(s, []byte("     ")...)
	}
	for i, j := len1-1, len(s)-1; i >= 0; {
		if s[i] <= '9' && s[i] >= '0' {
			for k := 0; k < len2; k++ {
				s[j-k] = insertElement[len2-k-1]
			}
			j -= 6
			i--
		} else {
			s[j] = s[i]
			i--
			j--
		}
	}
	fmt.Println(string(s))
}

func reverseWords(s string) string {
	str := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		} else {
			j := i
			for i >= 0 && s[i] != ' ' {
				i--
			}

			str = append(str, s[i+1:j+1]...)
			str = append(str, ' ')
		}
	}
	i := len(str) - 1
	for str[i] == ' ' {
		i--
	}
	return string(str[:i+1])
}

func reverseWords2(s string) string {
	str := []byte(s)
	slow := 0
	fast := 0
	for len(str) >= 0 && fast < len(str) && s[fast] == ' ' {
		fast++
	}
	for ; fast < len(s); fast++ {
		if fast > 0 && s[fast] == s[fast-1] && s[fast] == ' ' {
			continue
		}
		str[slow] = str[fast]
		slow++
	}
	if slow-1 > 0 && str[slow-1] == ' ' {
		slow--
	}

	reverse := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; {
			s[i] ^= s[j]
			s[j] ^= s[i]
			s[i] ^= s[j]
			i++
			j--
		}
	}

	reverse(str[:slow])

	start := 0
	for i := 0; i < slow; {
		for i < slow && str[i] != ' ' {
			i++
		}
		reverse(str[start:i])
		i++
		start = i
	}

	return string(str[:slow])
}

func rightRotate() {
	n := 0
	str := []byte{}
	fmt.Scanf("%d", &n)
	fmt.Scanln(&str)
	str = append(str[len(str)-n:], str[:len(str)-n]...)
	fmt.Println(string(str))
}

func rightRotate2() {
	n := 0
	str := []byte{}
	fmt.Scanln(&n)
	fmt.Scanln(&str)

	reverse := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; i++ {
			s[i] ^= s[j]
			s[j] ^= s[i]
			s[i] ^= s[j]
			j--
		}
	}

	reverse(str)
	reverse(str[:n])
	reverse(str[n:])
	fmt.Println(string(str))
}

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		sameN := 0
		for j, k := 0, i; j < len(needle); j++ {
			if k < len(haystack) && haystack[k] == needle[j] {
				sameN++
				k++
			} else {
				break
			}
		}
		if sameN == len(needle) {
			return i
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	// KMP算法实现
	getNext := func(str string) []int {
		j := 0
		next := make([]int, len(str))
		next[0] = 0
		for i := 1; i < len(str); i++ {
			for j > 0 && str[i] != str[j] {
				j = next[j-1]
			}
			if str[i] == str[j] {
				j++
			}
			next[i] = j
		}
		return next
	}

	next := getNext(needle)
	fmt.Println(next)

	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

func repeatedSubstringPattern(s string) bool {
	s2 := s[1:] + s[:len(s)-1]
	return strings.Contains(s2, s)
}

func repeatedSubstringPattern2(s string) bool {
	getNext := func(str string) []int {
		j := 0
		next := make([]int, len(str))
		next[0] = 0
		for i := 1; i < len(str); i++ {
			for j > 0 && str[i] != str[j] {
				j = next[j-1]
			}
			if str[i] == str[j] {
				j++
			}
			next[i] = j
		}
		return next
	}
	next := getNext(s)
	if next[len(s)-1] == 0 {
		return false
	}
	return len(s)%(len(s)-next[len(s)-1]) == 0
}

func main() {
	strStr2("mississippi", "issip")
}
