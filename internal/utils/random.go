package utils

import "math/rand"

func GenerateRandomIndexes(max, count int) []int {
	if count > max {
		count = max
	}

	randomIndexes := make([]int, count)
	for i := 0; i < count; i++ {
		randomIndexes[i] = rand.Intn(max)
	}

	return randomIndexes
}
