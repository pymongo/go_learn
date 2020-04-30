package tests

import (
	"fmt"
	"testing"
)

func uniquePaths(m int, n int) int {
	//m, n = m-1, n-1
	result := 1
	sum := m+n-2
	for i:=sum; i>n-1; i-- {
		result *= i
	}
	for i:=1; i<m; i++ {
		result /= i
	}
	return result
}

// faile(23, 12)
func TestUniquePaths(t *testing.T) {
	fmt.Println(uniquePaths(7, 3))
	//got := uniquePaths(1,1)
	//if got != 1 {
	//	t.Errorf("uniquePaths(m int, n int) = %d; want 1", got)
	//}
}
