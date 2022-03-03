package main

import "testing"

var (
	testArr1 = []int{1, 2, 3, 8, 9, 3, 2, 1}
	testArr2 = []int{1, 2, 1, 2, 2, 1}
	testArr3 = []int{7, 1, 2, 9, 7, 2, 1}
	val1     = 3
	val2     = 2
	val3     = 2
)

func TestFindMaxInSeriesNumberofSlice1(t *testing.T) {
	t.Logf("array : %v", testArr1)

	if findMaxInSeriesNumberofSlice(testArr1) != val1 {
		t.Errorf("FAIL! test case on array %v expected %d", testArr1, val1)
	}
}

func TestFindMaxInSeriesNumberofSlice2(t *testing.T) {
	t.Logf("array : %v", testArr2)
	if findMaxInSeriesNumberofSlice(testArr2) != val2 {
		t.Errorf("FAIL! test case on array %v expected %d", testArr2, val2)
	}
}

func TestFindMaxInSeriesNumberofSlice3(t *testing.T) {
	t.Logf("array : %v", testArr3)
	if findMaxInSeriesNumberofSlice(testArr3) != val3 {
		t.Errorf("FAIL! test case on array %v expected %d", testArr3, val3)
	}
}
