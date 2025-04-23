package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Print(squareArrayAndSort([]int{-4, -1, 0, 3, 10}))
	// fmt.Print(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	// fmt.Print(helixMatrix(3))
	// prefixSum()
	fmt.Println(diPrefixSum2())
}

// 螺旋矩阵
func generateMatrix(n int) [][]int {
	step := n - 1
	idx := 1
	turnTimes := 0
	dir := 0

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		matrix[0][i] = i + 1 // init first line
	}

	for m, i, j := n, 0, n-1; m <= n*n; m++ {
		matrix[i][j] = m
		if idx > step {
			turnTimes++
			dir = turnTimes % 4   // each time turns, direction changes
			if turnTimes%2 == 0 { // every two turns, step - 1
				step--
			}
			idx = 1
		}
		idx++
		switch dir {
		case 3: // 向右
			j++
		case 0: // 向下
			i++
		case 1: // 向左
			j--
		case 2: // 向上
			i--
		}
		fmt.Println(i, j)
	}
	return matrix
}

// 螺旋矩阵
func generateMatrix1(n int) [][]int {
	step := n
	idx := 1
	dir := 0

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for m, i, j := 1, 0, 0; m <= n*n; m++ {
		matrix[i][j] = m
		// next situation
		switch dir {
		case 0: // 向右
			j++
		case 1: // 向下
			i++
		case 2: // 向左
			j--
		case 3: // 向上
			i--
		}
		// next direction idx++
		if idx >= step {
			idx = 1
			dir++
			if dir >= 4 {
				// a round finish, begin next round
				dir = 0
				step -= 2
				i++
				j++
			}
		}
		fmt.Println(i, j)
	}
	return matrix
}

func intervalSum() {
	arrLen := 0
	fmt.Scanf("%d", &arrLen)
	arr := make([]int, arrLen)

	fmt.Scanf("%d", &arr[0])
	for i := 1; i < arrLen; i++ {
		fmt.Scanf("%d", &arr[i])
		arr[i] = arr[i] + arr[i-1]
	}

	a, b := 0, 0
	for {
		n, err := fmt.Scanf("%d %d", &a, &b)
		if n != 2 || err != nil {
			break
		}
		if a == 0 {
			fmt.Println(arr[b])
		} else {
			fmt.Println(arr[b] - arr[a-1])
		}
	}
}

func squareArrayAndSort(arr []int) []int {
	newArr := make([]int, len(arr))
	left, right := 0, len(arr)-1
	index := right
	for left <= right {
		if arr[left]*arr[left] >= arr[right]*arr[right] {
			newArr[index] = arr[left] * arr[left]
			left++
		} else {
			newArr[index] = arr[right] * arr[right]
			right--
		}
		index--
	}
	return newArr
}

func minSubArrayLen(s int, nums []int) int {
	l, r, sum, minLen := 0, 0, 0, len(nums)+1
	for ; r < len(nums); r++ {
		sum += nums[r]
		for sum >= s {
			if len := r - l + 1; len < minLen {
				minLen = len
			}
			sum -= nums[l]
			l++
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

func helixMatrix(n int) [][]int {
	sum := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	dir := 0
	for k, i, j := 1, 0, 0; k <= sum; k++ {
		matrix[i][j] = k
		switch dir {
		case 0: // 向右
			j++
		case 1: // 向下
			i++
		case 2: // 向左
			j--
		case 3: // 向上
			i--
		}
		if j == n || i == n || j == -1 || i == -1 || matrix[i][j] != 0 {
			switch dir {
			case 0:
				j--
				i++
			case 1:
				i--
				j--
			case 2:
				j++
				i--
			case 3:
				i++
				j++
			}
			dir = (dir + 1) % 4
		}
	}
	return matrix
}

func prefixSum() {
	arrLen := 0
	fmt.Scanf("%d\n", &arrLen)
	arr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		v := 0
		fmt.Scanf("%d\n", &v)
		if i == 0 {
			arr[i] = v
		} else {
			arr[i] = arr[i-1] + v
		}
	}
	fmt.Println(arr)
	for {
		a, b := 0, 0
		n, err := fmt.Scanf("%d %d\n", &a, &b)
		if n != 2 || err != nil {
			break
		}
		if a == 0 {
			fmt.Println(arr[b])
		} else {
			fmt.Println(arr[b] - arr[a-1])
		}
	}
}

func prefixSum2() {
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	arrLen, _ := strconv.Atoi(scaner.Text())
	arr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		scaner.Scan()
		v, _ := strconv.Atoi(scaner.Text())
		if i == 0 {
			arr[i] = v
		} else {
			arr[i] = arr[i-1] + v
		}
	}

	for scaner.Scan() {
		a, b := 0, 0
		fmt.Sscanf(scaner.Text(), "%d %d", &a, &b)
		if a == 0 {
			fmt.Println(arr[b])
		} else {
			fmt.Println(arr[b] - arr[a-1])
		}
	}
}

func diPrefixSum() int {
	x, y := 0, 0
	fmt.Scanf("%d %d\n", &x, &y)

	arr1 := make([]int, x)
	arr2 := make([]int, y)
	sum := 0
	min := 0

	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			v := 0
			fmt.Scan(&v)
			arr1[i] += v
			arr2[j] += v
			sum += v
		}
	}

	min = sum

	for i := 0; i < x; i++ {
		if i != 0 {
			arr1[i] = arr1[i] + arr1[i-1]
		}
		diff1 := abs(sum - 2*arr1[i])
		if diff1 < min {
			min = diff1
		}
	}

	for i := 0; i < y; i++ {
		if i != 0 {
			arr2[i] = arr2[i] + arr2[i-1]
		}
		diff2 := abs(sum - 2*arr2[i])
		if diff2 < min {
			min = diff2
		}
	}

	return min
}

func diPrefixSum2() int {
	x, y := 0, 0
	fmt.Scanf("%d %d\n", &x, &y)

	m := make([][]int, x+1)
	m[0] = make([]int, y+1)
	for i := 1; i <= x; i++ {
		m[i] = make([]int, y+1)
		for j := 1; j <= y; j++ {
			v := 0
			fmt.Scan(&v)
			m[i][j] = m[i-1][j] + m[i][j-1] - m[i-1][j-1] + v
		}
	}

	min := m[x][y]

	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	for i := 1; i <= x; i++ {
		diff1 := abs(m[x][y] - 2*m[i][y])
		if diff1 < min {
			min = diff1
		}
	}

	for i := 1; i <= y; i++ {
		diff1 := abs(m[x][y] - 2*m[x][i])
		if diff1 < min {
			min = diff1
		}
	}

	return min
}
