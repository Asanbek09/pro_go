package main

import (
	"testing"
	"sort"
)

func TestSum(t *testing.T) {
	testValues := []int {10, 20, 30}
	_, sum := sortAndTotal(testValues)
	expected := 60
	if (sum != expected) {
		t.Fatalf("Expected %v, Got %v", expected, sum)
	}
}

func TestSort(t *testing.T) {
	testValues := []int {1, 23, 49, 93, 2, 5, 43, 9}
	sorted, _ := sortAndTotal(testValues)
	if (!sort.IntsAreSorted(sorted)) {
		t.Fatalf("Unsorted data: %v", sorted)
	}
}