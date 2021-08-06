package main

import (
	"fmt"
	"strings"
)

/*
Complete the method/function so that it converts dash/underscore
delimited words into camel casing. The first word within the
output should be capitalized only if the original word was capitalized
(known as Upper Camel Case, also often referred to as Pascal case).

Examples
"the-stealth-warrior" gets converted to "theStealthWarrior"
"The_Stealth_Warrior" gets converted to "TheStealthWarrior"
 */

func ToCamelCase(s string) string {
	splitted:=[]string{}
	if strings.Contains(s, "-") {
		splitted=strings.Split(s, "-")
	} else {
		splitted=strings.Split(s, "_")
	}
	fmt.Println(splitted)
	arr:=make([]string, 4)
	for i,_:= range splitted{
		strings.ToTitle(arr[i])
	}
	//fmt.Println(arr)
	//fmt.Sprintf()
	//fmt.Sprint()
	//fmt.Fprintln()
	//fmt.Printf()
	//fmt.Fprintln()
	//fmt.Fprint()
	// your code
	return ""
}

func main() {
	s:=("the-stealth-warrior")
	s1:=("The_Stealth_Warrior")
	fmt.Println(ToCamelCase(s))
	fmt.Println(ToCamelCase(s1))
}