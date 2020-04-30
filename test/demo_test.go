package test

import "testing"

func Abs(num int32) int32 {
	return num
}

// Files containing tests should with the `_test` suffix
func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
