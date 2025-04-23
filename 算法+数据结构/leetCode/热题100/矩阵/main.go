package main

import "fmt"

// 1. 矩阵置零

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rToZ, cToZ := make(map[int]bool), make(map[int]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rToZ[i], cToZ[j] = true, true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rToZ[i] || cToZ[j] {
				matrix[i][j] = 0
			}
		}
	}

	// for r, _ := range rToZ {
	// 	for c:=0; c<n; c++ {
	// 		matrix[r][c] = 0
	// 	}
	// }

	// for c, _ := range cToZ {
	// 	for r:=0; r<m; r++ {
	// 		matrix[r][c] = 0
	// 	}
	// }
}

func setZeroes2(matrix [][]int) {
	// 用第一列和第一行做为标记，且只用一个标记变量来记录第一列是否原本就存在0，第一行是否原本存在0已经记录在了[0][0]位置
	// 注意从最后一行开始倒序处理，防止第一行的标记被清空
	m, n := len(matrix), len(matrix[0])
	col0 := false
	for r := 0; r < m; r++ {
		if matrix[r][0] == 0 {
			col0 = true
		}
		for c := 1; c < n; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				matrix[r][0] = 0
			}
		}
	}

	for r := m - 1; r >= 0; r-- {
		for c := 1; c < n; c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
		if col0 {
			matrix[r][0] = 0
		}
	}
}

// 2. 螺旋矩阵

func spiralOrder(matrix [][]int) []int {
	rB, bB, lB, uB := len(matrix[0])-1, len(matrix)-1, 0, 0
	idx, sum, dir, r, c := 0, (rB+1)*(bB+1), 0, 0, 0
	res := make([]int, sum)
	for ; idx != sum; dir = (dir + 1) % 4 {
		switch dir {
		case 0:
			for ; c <= rB; c++ {
				res[idx] = matrix[r][c]
				idx++
			}
			r++
			c--
		case 1:
			for ; r <= bB; r++ {
				res[idx] = matrix[r][c]
				idx++
			}
			r--
			c--
		case 2:
			for ; c >= lB; c-- {
				res[idx] = matrix[r][c]
				idx++
			}
			r--
			c++
		case 3:
			for ; r > uB; r-- {
				res[idx] = matrix[r][c]
				idx++
			}
			r++
			c++
			// 更新边界
			rB, bB, lB, uB = rB-1, bB-1, lB+1, uB+1
		}
	}
	return res
}

func main() {
	fmt.Println(`1. 矩阵置零: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]`)
	matrix1 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes2(matrix1)
	fmt.Println(matrix1)

	fmt.Println(`2. 螺旋矩阵: matrix = [[1,2,3,4,5],[6,7,8,9,10],[11,12,13,14,15],[16,17,18,19,20]]`)
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}}))
}
