package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("\nResult -> %v", EditWord("telkom", "telecom"))
	fmt.Printf("\nResult -> %v\n", EditWord("telkom", "tlkom"))
}

func EditWord(firstWord, secondWord string) bool {
	var result []string
	var temp string

	if len(secondWord) < len(firstWord) || len(firstWord) < len(secondWord) {
		temp = secondWord
		secondWord = firstWord
		firstWord = temp

	} else if len(firstWord) == len(secondWord) {
		r, _ := regexp.Compile("[" + firstWord + "]")
		result = r.FindAllString(secondWord, 2)
		join := strings.Join(result, "")
		var result2 []int
		var t int

		for i := 0; i < len(join); i++ {
			if string(join[i]) != " " {
				t++
				result2 = append(result2, t)
			} else if i == len(join)-1 {
				result2 = append(result2, t)
			}
		}

		if len(result2) == len(secondWord)-1 {
			return true
		}
		return false
	}

	r, _ := regexp.Compile("[" + secondWord + "]")
	result = r.FindAllString(firstWord, -1)
	join := strings.Join(result, "")

	if len(join) == len(firstWord) {
		return true
	} else {
		return false
	}
}
