package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	output := anagramFinder(input)
	fmt.Printf("%+v\n", output)
}

type anaword struct {
	Val       string
	SortedVal string
}

func anagramFinder(input []string) [][]string {
	// populate anaword
	words := make([]anaword, 0)
	for _, v := range input {
		words = append(words, anaword{
			Val:       v,
			SortedVal: sortWord(v),
		})
	}

	// compare anaword
	wordMap := make(map[string][]string, 0)
	for _, v := range words {
		if _, ok := wordMap[v.SortedVal]; !ok {
			wordMap[v.SortedVal] = make([]string, 0)
		}
		wordMap[v.SortedVal] = append(wordMap[v.SortedVal], v.Val)
	}

	// transform from map to slice
	res := make([][]string, 0)
	for _, v := range wordMap {
		res = append(res, v)
	}

	return res
}

func sortWord(word string) string {
	wordRune := []rune(word)
	sort.Slice(wordRune, func(i, j int) bool { return wordRune[i] < wordRune[j] })
	return string(wordRune)
}
