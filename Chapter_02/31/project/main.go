package main

import (
	"sort"
	//"fmt"
	"log"
)

func sortAndTotal(vals []int) (sorted []int, total int) {
	var logger = log.New(log.Writer(), "sortAndTotal: ", log.Flags() | log.Lmsgprefix)
	logger.Printf("Invoked with %v values", len(vals))
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	logger.Printf("Sorted data: %v", sorted)
	for _, val := range sorted {
		total += val
		//total++
	}
	logger.Printf("Total: %v", total)
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