package main

import (
	"fmt"
)

func replaceAllSpaces(a string) string {
	runes := []rune(a)
	last_position := get_last_position(runes)
	count_space := count_space(runes, last_position+1)
	new_length := last_position + 1 + 2*count_space

	new_runes := make([]rune, new_length)

	j := 0
	for i := 0; i <= last_position; i++ {
		if runes[i] == ' ' {
			new_runes[j] = '%'
			new_runes[j+1] = '2'
			new_runes[j+2] = '0'
			j += 3
		} else {
			new_runes[j] = runes[i]
			j++
		}
	}
	return string(new_runes)
}

func get_last_position(runes []rune) int {
	for i := len(runes) - 1; i > -1; i-- {
		if runes[i] != ' ' {
			return i
		}
	}
	return 0
}
func count_space(runes []rune, last_position int) int {
	count := 0
	for i := 0; i < last_position; i++ {
		if runes[i] == ' ' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(replaceAllSpaces("this is a space  "))
	fmt.Println(replaceAllSpaces("this is another   space"))
}
