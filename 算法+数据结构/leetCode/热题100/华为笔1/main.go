package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func patchUpdate() {
	n := 0
	fmt.Scanln(&n)
	m := map[string][]string{}
	for i := 0; i < n; i++ {
		str1, str2 := "", ""
		fmt.Scanln(&str1, &str2)
		if _, ok := m[str1]; ok {
			m[str1] = append(m[str1], str2)
		} else {
			m[str1] = []string{str2}
		}
	}

	m2 := map[string]int{}
	var dfs func(k string) int
	dfs = func(k string) int {
		maxDepth := 0
		for _, s := range m[k] {
			d := 0
			if _, ok := m2[s]; ok {
				d = m2[s]
			} else {
				d = dfs(s)
				m2[s] = d
			}
			maxDepth = max(maxDepth, d+1)
		}
		return maxDepth
	}

	res := []string{}
	maxDepth := 0
	for k := range m {
		d := 0
		if _, ok := m2[k]; ok {
			d = m2[k]
		} else {
			d = dfs(k)
		}
		if maxDepth < d {
			maxDepth = d
			res = []string{k}
		} else if maxDepth == d {
			res = append(res, k)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
			if res[i][k] < res[j][k] {
				return true
			}
		}
		if len(res[i]) < len(res[j]) {
			return true
		} else {
			return false
		}
	})

	fmt.Println(res)
}

// 4
// a d
// a b 1
// a c 3
// a d 6
// b c 1
// b d 4
// c d 1
// 0000

func shortestPath() {
	n := 0
	fmt.Scanln(&n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	start, end := text[0], text[2]
	m := map[byte]map[byte]int{}
	for scanner.Scan() {
		str := scanner.Text()
		if str == "0000" {
			break
		}
		sT, _ := strconv.Atoi(str[4:])
		if _, ok := m[str[0]]; ok {
			m[str[0]][str[2]] = sT
		} else {
			m[str[0]] = map[byte]int{str[2]: sT}
		}
		if _, ok := m[str[2]]; ok {
			m[str[2]][str[0]] = sT
		} else {
			m[str[2]] = map[byte]int{str[0]: sT}
		}
	}

	fmt.Println(n, start, end, m)

	sL := math.MaxInt
	path := []byte{}
	pass := [26]int{}
	var dfs func(start, end byte, l int) bool
	dfs = func(start, end byte, l int) bool {
		if start == end {
			fmt.Println("to end")
			if l < sL {
				path = []byte{}
				sL = l
				return true
			}
			return false
		}

		// 经过就标记
		pass[start-'a'] = 1

		flag := false
		for k, v := range m[start] {
			if pass[k-'a'] == 1 {
				continue
			}
			fmt.Printf("dfs %c -> %c: %d\n", start, k, v)
			if dfs(k, end, l+v) {
				flag = true
				path = append(path, k)
			}
		}

		// 结束就取消
		pass[start-'a'] = 0

		return flag
	}

	if dfs(start, end, 0) {
		path = append(path, start)
	}

	fmt.Println(path)
}

func rotateMatrix(mtx [][]int) int {
	m, n := len(mtx), len(mtx[0])
	arr := make([]int, m*n)
	dir := 0
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	l, r, u, b := 0, n-1, 0, m-1
	for i, x, y := 0, 0, 0; i < m*n; {
		// 先赋值当前位置，再探测下一步是否合法，不合法就做方向调整，然后进入合法的下一步
		arr[i] = mtx[y][x]
		nx, ny := x+dx[dir], y+dy[dir]
		if nx < l || nx > r || ny > b || (dir != 3 && ny < u) || (dir == 3 && ny < u+1) { // 注意绕完一圈时最后的判断要是ny<u+1
			dir = (dir + 1) % 4
			if dir == 0 {
				l, r, u, b = l+1, r-1, u+1, b-1
			}
		}
		x, y = x+dx[dir], y+dy[dir]
		i++
	}

	return mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) int {
	if l>=r {
		return 0
	}
	mid := l+(r-l)/2

	res := mergeSort(arr, l, mid) + mergeSort(arr, mid+1, r)

	newA := []int{}
	i,j:=l,mid+1
	for i<=mid&&j<=r {
		if arr[i] <= arr[j] {
			res += j-(mid+1)
			newA = append(newA, arr[i])
			i++
		} else {
			newA = append(newA, arr[j])
			j++
		}
	}

	for ; i<=mid; i++ {
		res += r-(mid+1)+1
		newA = append(newA, arr[i])
	} 
	for ; j<=r; j++ {
		newA = append(newA, arr[j])
	}

	for i:=l; i<=r; i++ {
		arr[i] = newA[i-l]
	}

	return res
}

func main2() {
	// patchUpdate()
	// shortestPath()
	// fmt.Println(rotateMatrix([][]int{
	// 	{3, 2, 1},
	// 	{6, 5, 4},
	// 	{9, 8, 7},
	// }))


}
