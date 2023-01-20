package main

import "fmt"

func main() {
	nextEvenNumber := newEvenNumbersGenerator(19)
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
}

func newEvenNumbersGenerator(maximum uint) func() (uint, bool) {
	i := uint(0)
	return func() (result uint, ok bool) {
		result = i
		if result <= maximum {
			i += 2
			return result, true
		}
		return 0, false

	}
}
