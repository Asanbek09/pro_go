package main

import (
	"sort"
	"fmt"
)

func sortAndTotal(vals []int) (sorted []int, total int) {
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	for _, val := range sorted {
		total += val
		total++
	}
	return
}

func main() {
	nums := []int {90, 56, 82, 19, 32, 74}
	sorted, total := sortAndTotal(nums)
	fmt.Println("Sorted data:", sorted)
	fmt.Println("Total:", total)
}