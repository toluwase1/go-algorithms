package main

import "fmt"

func main() {
	fmt.Println(hackerrankInString("hackerrank"))
}

func hackerrankInString(s string) string {
	// Write your code here
	required:= "hackerrank"
	//length:=int32(len(required))

	count:=0
	for i:=range s{
		if s[i]==required[count] {
			count+=1
		}
		//len:=int(length)
		if count==10 {
			return "YES"
		}
	}
	return "NO"
}

func hackerrankInString11(s string) string {
	// Write your code here
	required:= "hackerrank"
	length:=int32(len(required))
	count:=0
	for j:=range s{
		if s[j] == required[count] {
			count+=1
		}
		len:=int(length)
		if count==len {
			return "YES"
		}
	}
	return "NO"
}




func hackerrankInStringabc(s string) string {
	idx := 0
	str := "hackerrank"
	for i:= range s{
		if s[i] == str[idx]{
			idx++
		}
		if idx == len(str){
			return "YES"
		}
	}
	return "NO"
}




