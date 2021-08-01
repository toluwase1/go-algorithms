package main

import "fmt"

//loop through
//check if even or odd
func main() {
	slices:=[]int{1,2,4,6,8}
	//FindOutlier(slices)
	fmt.Println(FindOutlier(slices))
}

func FindOutlier(integers [] int) int {
	var holdEven []int
	var holdOdd []int
	for _, v:= range integers {

		if v%2==0 {
			//holdEven = append(holdEven, integers[v])
			holdEven = append(holdEven, v)
		} else {
			holdOdd= append(holdOdd, v)
		}
	}
	if len(holdOdd)==1 {
		return holdOdd[0]
	}

	return holdEven[0]
}

//{{1,2,3}, {1,2,4} , {1,2,3}}