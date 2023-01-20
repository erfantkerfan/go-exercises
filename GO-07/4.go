package main

import "fmt"

func main() {
	salt := "->"
	fmt.Println(joinStrings(salt))
	fmt.Println(joinStrings(salt, "60"))
	fmt.Println(joinStrings(salt, "60", "15"))
	fmt.Println(joinStrings(salt, "60", "15", "5H17"))

	foo := []string{"60", "15", "r3411Y", "5H17"}
	fmt.Println(joinStrings(salt, foo...))
}

func joinStrings(joiner string, strings ...string) (result string) {
	result = ""
	for k, str := range strings {
		if k == 0 {
			result += str
			continue
		}
		result += joiner + str
	}
	return result
}
