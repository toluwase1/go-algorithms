package main

import "fmt"

func main() {
	n := 123454321
	r := reverse(n)
	fmt.Println(r == n)
}

func reverse(n int) int {
	r := 0
	for {
		if n > 0 {
			r = r*10 + n%10
			n = n / 10
		} else {
			break
		}
	}
	return r
}