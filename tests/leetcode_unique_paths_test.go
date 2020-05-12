package tests

import (
	"fmt"
	"testing"
)

// leetcode62. unique-paths
func uniquePaths(m int, n int) int {
	// 端点是4x4，但是棋盘的格子就3x3
	max, min := m-1, n-1
	if min > max {
		max, min = min, max
	}
	result := 1
	sum := max+min
	for i:=0; i<min; i++ {
		result *= sum-i
		result /= i+1
	}
	return result
}

// faile(23, 12)
func TestUniquePaths(t *testing.T) {
	fmt.Println(uniquePaths(23, 12))
	//got := uniquePaths(1,1)
	//if got != 1 {
	//	t.Errorf("uniquePaths(m int, n int) = %d; want 1", got)
	//}
}
