package tests

import (
	"fmt"
	"testing"
)

func uniquePaths(m int, n int) int {
	// C(m,m+n)
	// C(3,5) 5*4*3 / 1*2*3
	result := 1
	sum := m+n
	for i:=sum; i>n; i-- {
		result *= i
	}
	for i:=1; i<m; i++ {
		result /= i
	}
	return result
}

func TestUniquePaths(t *testing.T) {
	fmt.Println(uniquePaths(2, 3))
	//got := uniquePaths(1,1)
	//if got != 1 {
	//	t.Errorf("uniquePaths(m int, n int) = %d; want 1", got)
	//}
}
