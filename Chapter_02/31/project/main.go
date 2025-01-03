package main

import (
	"sort"
	//"fmt"
	"log"
)

func sortAndTotal(vals []int) (sorted []int, total int) {
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	for _, val := range sorted {
		total += val
		//total++
	}
	return
}

func main() {
	nums := []int {90, 56, 82, 19, 32, 74}
	sorted, total := sortAndTotal(nums)
	log.Print("Sorted data:", sorted)
	log.Print("Total:", total)
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}