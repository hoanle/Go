package question11

import (
	"fmt"
)

func isUnique(a string) bool {
	if len(a) > 128 {
		return false
	} else {
		var char_list [128]bool
		for _, val := range a {
			if char_list[val] {
				return false
			}
			char_list[val] = true
		}
		return true
	}
}

func main() {
	fmt.Println(isUnique("abcdef1234567s "))
	fmt.Println(isUnique("abcdef1234567s1"))
}
