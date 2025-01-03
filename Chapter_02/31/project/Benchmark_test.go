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
	b.ResetTimer()
	for i := 0; i < size; i++ {
		b.StopTimer()
		data[i] = rand.Int()
	}
	b.StartTimer()
	sortAndTotal(data)
}