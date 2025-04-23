package main

import (
	"fmt"
	// "sort"
)

func partitionLabels(s string) []int {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}

	res := []int{}
	start, end := 0, 0
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			res = append(res, end-start+1)
			start = end+1
		}
	}
	
	return res
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
