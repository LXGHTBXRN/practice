package main

import "fmt"

func longestString(strings ...string) string {
	var longest string

	for _, str := range strings {
		if len(str) > len(longest) {
			longest = str
		}
	}
	return longest
}

func main() {
	result := longestString("cat", "frog", "car", "elephant", "plane")
	fmt.Println("Самая длиная строка", result)
}
