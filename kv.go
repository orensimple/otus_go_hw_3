package main

import (
	"fmt"
	"sort"
)

type Kv struct {
	Key   string
	Value int
}

func MapToSlice(wordsDic map[string]int) []Kv {
	var resultSlice []Kv
	for k, v := range wordsDic {
		resultSlice = append(resultSlice, Kv{k, v})
	}
	return resultSlice
}

func PrintSort10(resultSlice []Kv) {
	sort.Slice(resultSlice, func(i, j int) bool {
		return resultSlice[i].Value > resultSlice[j].Value
	})
	for i := 0; i < 10 && i < len(resultSlice); i++ {
		fmt.Printf("%s - %d\n", resultSlice[i].Key, resultSlice[i].Value)
	}
}
