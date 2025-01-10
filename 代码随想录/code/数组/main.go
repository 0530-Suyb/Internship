package main

import (
	"fmt"
)

func main() {
	generateMatrix1(3)
}

func generateMatrix(n int) [][]int {
    step := n-1
    idx := 1
    turnTimes := 0
    dir := 0

    matrix := make([][]int, n)
    for i:=0; i<n; i++ {
        matrix[i] = make([]int, n)
        matrix[0][i] = i+1 // init first line
    }

    for m,i,j:=n,0,n-1; m<=n*n; m++ {
        matrix[i][j] = m
        if idx > step {
            turnTimes++
            dir = turnTimes % 4 // each time turns, direction changes
            if turnTimes % 2 == 0 {   // every two turns, step - 1 
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
        fmt.Println(i,j)
    }
    return matrix
}

func generateMatrix1(n int) [][]int {
    step := n
    idx := 1
    dir := 0

    matrix := make([][]int, n)
    for i:=0; i<n; i++ {
        matrix[i] = make([]int, n)
    }

    for m,i,j:=1,0,0; m<=n*n; m++ {
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
		// next direction o                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        n next situation
		idx++
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
        fmt.Println(i,j)
    }
    return matrix
}

func intervalSum() {
	arrLen := 0
	fmt.Scanf("%d", &arrLen)
	arr := make([]int, arrLen)

	fmt.Scanf("%d", &arr[0])
	for i:=1; i<arrLen; i++ {
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