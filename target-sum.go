package main

import "fmt"

//import "strconv"
func TwoNumberSum(array []int, target int) []int {
	//for range array{
	//
	//}
	for i,_:= range array {
		firstNum:=array[i]
		for _,j:=range array{
			secondNum:=array[j]
			if firstNum + secondNum==target {
				return [] int {firstNum, secondNum}
			}
		}
	}
	return array
}

func reverseSlice(a []int) [] int {
	l := len(a)
	for i := 0; i < l/2; i ++ {
		x := a[l-i-1]
		a[l-i-1] = a[i]
		a[i] = x
	}
	return a
}
func main() {
//var array int[] = {3,5,-4,8,11,1,-1,6}
//va
//	var (
//	target = 10
//	)
	// use break/continue with label on outer loop
//here:
//	for i := 0; i < 2; i++ {
//		for j := i + 1; j < 3; j++ {
//			if i == 0 {
//				continue here
//			}
//			fmt.Println(j)
//			if j == 2 {
//				break
//			}
//		}
//	}


	//make(types.Array, [3,5,-4,8,11,1,-1,6], )
	a := [] int {3,5,-4,8,11,1,-1,6}
	ab:=[] int {}
	for i:= len(a)-1; i>=0; i-- {

		ab = append(ab, a[i])
	}
	fmt.Print("reversed1: ",ab)

	for i:=0; i< len(a)/2; i++ {
		temp:= a[i]
		a[i] =a[len(a)-i]
		a[len(a)-i]= temp
	}
	fmt.Print("reversed1: ",)

	//fmt.Println(TwoNumberSum(a, 10))
	//fmt.Println(reverseSlice(a))
	////fmt.Println(multiplicationTable)
	//for i:=1; i<=10; i++{
	//	for j:=1; j<10; j++ {
	//		fmt.Print(i*j, " ")
	//	}
	//	fmt.Println();
	//}
}
func multiplicationTable()  {
	for i:=1; i<=10; i++{
		for j:=1; j<10; j++ {
			fmt.Print(i*j, " ")
		}
		fmt.Println();
	}
}



//hash_map = {}
//for i in range(len(nums)):
//if nums[i] in hash_map:
//return [i, hash_map[nums[i]]]
//else:
//hash_map[target - nums[i]] = i

//func TwoSums(numbers []int, target int) [2]int {
////	strconv.Itoa(numbers)
//	maps:=map[int]int{}
//	for i,j:= range numbers{
//		if numbers[i]+numbers[i+1]==target {
//
//		}
//	}
//	return [2]int{}
//}

