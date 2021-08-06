package main

import (
	"fmt"
	"strings"
	"unicode"
)

//func DuplicateEncode(word string) string {
//	//loop through the string
//	//find a recurring value, give it ")"
//	//else, give it (
//	saveWord:=""
//	for i,j:= range word{
//		if word[i]==word[i+1] {
//			b := saveWord == word[i]
//			i++
//		}
//
//	}
//
//	return // ...
//}

func EncodeDuplicateParens(input string) string {
	encoded := strings.Builder{}
	occurrences := make(map[rune]int)

	for _, character := range input {
		occurrences[unicode.ToLower(character)]++
	}

	for _, character := range input {
		if occurrences[unicode.ToLower(character)] > 1 {
			encoded.WriteRune(')')
		} else {
			encoded.WriteRune('(')
		}
	}
	return encoded.String()
}

func Encoded(word string) string {

	withDuplicate:=strings.Builder{}
	mapData:= make(map[rune]int)

	for _,j:= range word{
		mapData[unicode.ToLower(j)]++
	}
	for _,j:=range word{
		if mapData[unicode.ToLower(j)]>1 {
			withDuplicate.WriteRune(')')
		} else {
			withDuplicate.WriteRune('(')
		}
	}
	return withDuplicate.String()
}





func DuplicateEncode1(word string) string {
	word = strings.ToLower(word)
	mapWord := map[rune]int{}
	for _, v := range word {
		mapWord[v]=mapWord[v]+1
		fmt.Println(mapWord)
	}
	final:=""
	for _, c := range word {
		if mapWord[c] == 1 {
			final += "("
		} else {
			final += ")"
		}
	}
	fmt.Println(final)
	return final
}

func main() {
	fmt.Println(DuplicateEncode1("abbde"))
	var array = [] int {1,2,3,4,5}
	fmt.Println(sumArray(array))
	fmt.Println(findNum(array))
}

func sumArray(arr [] int) int {
	sum:=0
	for i,_:=range arr{
		sum = sum+ arr[i]
	}
	return sum
}

func findNum(arr [] int) int {
	a:=3
	for i,_:= range arr{
		if arr[i]==a {
			return a
		}

	}
	return -1
}