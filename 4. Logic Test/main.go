package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortWord(word string) string {
	arrWord := strings.Split(word, "")
	sort.Strings(arrWord)
	return strings.Join(arrWord, "")
}

func main() {
	mapAnagrams := make(map[string][]string, 0)
	arrStrings := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	for _, str := range arrStrings {
		sortedWord := sortWord(str)
		mapAnagrams[sortedWord] = append(mapAnagrams[sortedWord], str)
	}

	for _, anagrams := range mapAnagrams {
		fmt.Printf("%+v \n", anagrams)
	}
}
