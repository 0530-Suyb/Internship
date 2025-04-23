package main

import "fmt"

func huiwen(s string) {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	// 不能如此，因为f[i][j]的值是依赖于f[i+1][j-1]的值的，所以i得从大到小，j得从小到大
	// for i := 0; i < n; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		f[i][j] = s[i] == s[j] && f[i+1][j-1]
	// 	}
	// }

	fmt.Println(f)
}

func main() {
	fmt.Println((-6)%5)
}
