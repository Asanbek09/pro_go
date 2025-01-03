package main

import (
	"testing"
	"math/rand"
	"time"
)

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 250
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	sortAndTotal(data)
}