package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(min(1, 2))
	fmt.Println(min(-200, -100, 0, 100, 200))
	// fmt.Println(min(1))
	fmt.Println(max(1, 2))
	fmt.Println(max(-200, -100, 0, 100, 200))
	// fmt.Println(max(1))
}

func min(a, b float64, others ...float64) (result float64) {
	result = math.Min(a, b)
	for _, c := range others {
		result = math.Min(result, c)
	}
	return result
}

func max(a, b float64, others ...float64) (result float64) {
	result = math.Max(a, b)
	for _, c := range others {
		result = math.Max(result, c)
	}
	return result
}
