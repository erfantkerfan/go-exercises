package main

import "fmt"

func main() {
	fmt.Println(mul(4, 2))
	fmt.Println(mul(8, 16))
	fmt.Println(mul(-10, 10))
	fmt.Println(mul(-16, 8))
	fmt.Println(mul(-16, 9))
}

func mul(a, b int8) (c int8, ok bool) {
	c = a * b
	if a == 0 || b == 0 {
		return c, true
	} else if c/a == b {
		return c, true
	}
	return 0, false
}
