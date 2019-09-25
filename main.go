package main

import (
	"strings"
)

func main() {
	input := `Иванов Иван Иванович Иван Иванович Иван Иванов 1, 1 1 1 жил был дядя фдор а у него жил был дядя и тетя`
	replaceText := strings.ReplaceAll(input, ",", "")
	wordsSlice := strings.Split(replaceText, " ")
	wordsDic := map[string]int{}
	for _, word := range wordsSlice {
		wordsDic[word] = wordsDic[word] + 1
	}
	result := MapToSlice(wordsDic)
	PrintSort10(result)
}
