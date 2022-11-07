package helper

import "testing"

func TestSum_Success(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expect := 15

	actual := Sum(nums...)

	if actual != expect {
		t.Errorf("fail! expect : %d got : %d", expect, actual)
	}
}

func TestSum_Fail(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expect := 14

	actual := Sum(nums...)

	if actual == expect {
		t.Errorf("fail! expect : %d got : %d", expect, actual)
	}
}
