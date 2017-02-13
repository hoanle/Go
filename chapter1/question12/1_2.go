package main

import (
	"fmt"
)

func isPermutation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	} else {
		var char_list [128]int
		for _, c := range a {
			char_list[c] += 1
		}
		for _, c := range b {
			if char_list[c] == 0 {
				return false
			}
			char_list[c] -= 1
		}

		return true
	}
}
func main() {
	fmt.Println(isPermutation("abc", "cab"))
	fmt.Println(isPermutation("abcd", "cabc"))
	fmt.Println(isPermutation("abcd", "acdb"))
}
