package main

import (
	"testing"
)

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestMerge(t *testing.T) {
	expected_result := []int{1, 2, 3, 4, 5, 6, 7, 8}

	slice := []int{3, 2, 1, 4, 8, 6, 5, 7}
	left_slice := []int{1, 2, 3, 4}
	right_slice := []int{5, 6, 7, 8}

	Merge(left_slice, right_slice, &slice)
	ok := testEq(slice, expected_result)
	if ok == false {
		t.Errorf("FAIL: %v", slice)
	} else if ok == true {
		t.Logf("PASSED: %v", slice)
	}

}
