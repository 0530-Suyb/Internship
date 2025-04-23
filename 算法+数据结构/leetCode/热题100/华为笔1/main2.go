package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 将IP地址转换为整数
func ipToInt(ip string) int {
	parts := strings.Split(ip, ".")
	result := 0
	for i := 0; i < 4; i++ {
		num, _ := strconv.Atoi(parts[i])
		result = result<<8 + num
	}
	return result
}

// 最大不重叠业务数量
func maxNonOverlappingBusinesses(yw [][2]string) int {
	// 将IP范围转换为整数范围
	intRanges := make([][2]int, len(yw))
	for i, r := range yw {
		intRanges[i] = [2]int{ipToInt(r[0]), ipToInt(r[1])}
	}

	// 按业务的结束空间排序
	sort.Slice(intRanges, func(i, j int) bool {
		return intRanges[i][1] < intRanges[j][1]
	})

	count := 0
	end := -1 // 记录上一个选择的业务的结束空间

	for _, business := range intRanges {
		// 如果当前业务的起始空间大于等于上一个业务的结束空间
		if business[0] >= end {
			count++
			end = business[1] // 更新结束空间
		}
	}

	return count
}

func main() {
	// 示例业务占用的IP范围
	yw := [][2]string{
		{"172.168.1.1", "172.168.1.10"},
		{"172.168.1.2", "172.168.1.3"},
		{"172.168.1.4", "172.168.1.5"},
		{"172.168.1.6", "172.168.1.7"},
	}

	// 计算最大满足的业务数量
	result := maxNonOverlappingBusinesses(yw)
	fmt.Println("最大满足的业务数量:", result)
}
