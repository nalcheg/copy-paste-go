package main

import (
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestB(t *testing.T) {
	arr := make([]int, 10000)

	for k, _ := range arr {
		arr[k] = gofakeit.Number(0, 100000)
	}

	sorted := BubbleSort(arr)

	_ = sorted
}
