package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := `Иванов Иван Иванович Иван Иванович Иван Иванов 1, 1 1 1 жил был дядя фдор а у него жил был дядя и тетя`
	lowerText := strings.ToLower(input)
	replaceText1 := strings.ReplaceAll(lowerText, ",", "")
	wordsSlice := strings.Split(replaceText1, " ")
	wordsDic := map[string]int{}
	for _, word := range wordsSlice {
		wordsDic[word] = wordsDic[word] + 1
	}

	type kv struct {
		Key   string
		Value int
	}

	var resultSlice []kv
	for k, v := range wordsDic {
		resultSlice = append(resultSlice, kv{k, v})
	}

	sort.Slice(resultSlice, func(i, j int) bool {
		return resultSlice[i].Value > resultSlice[j].Value
	})

	for i := 0; i < 10 && i < len(resultSlice); i++ {
		fmt.Printf("%s, %d\n", resultSlice[i].Key, resultSlice[i].Value)
	}
}
