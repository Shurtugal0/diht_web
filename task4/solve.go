package main

import (
	"unicode"
	"strings"
)

func RemoveEven(slice []int)(out_slice [] int) {
	for _, num := range(slice) {
		if num % 2 == 1 {
			out_slice = append(out_slice, num)
		}
	}
	return
}

func PowerGenerator(number int) (func() int) {
	pow := 1
	return func() (result int) {
		result = 1
		for i := 0; i < pow; i++ {
			result *= number
		}
		pow++
		return
	}
}

func DifferentWordsCount(str string) int {
	diff_words := make(map[string]bool)
	word := ""
	for _, char := range(str) {
		if unicode.IsLetter(char) {
			word = word + string(char)
			word = strings.ToLower(word)
		} else {
			if word != "" {
				diff_words[word] = true
			}
			word = ""
		}
	}
	if word != "" {
		diff_words[word] = true;
	}
	return len(diff_words)
}