package main

import "fmt"

// concurrency vs parallelism
//rob pike simplicity is complicated
//go doc
//educative.io
// https://invite.slack.golangbridge.org/
//blog
//buy
//ebook
func main() {
	slices1:=[]int32{1,2,1,2,1,3,2}
	fmt.Println(sockMerchant(7,slices1))
}
func BubbleSort(array[] int32)[]int32 {
	for i:=0; i< len(array)-1; i++ {
		for j:=0; j < len(array)-i-1; j++ {
			if (array[j] > array[j+1]) {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
func sockMerchant(n int32, ar []int32) int32 {
	count:=int32(0)
	BubbleSort(ar)
	for i:=1; i< len(ar); i++ {
		if ar[i]==ar[i-1] {
			i++
			count++
		}
	}
	return count
}

// Write your code here
//sort the array
//count the pairs in the array.
//return count
//a:=int(n)

//for i,_:= range ar {
//	if ar[i]==ar[i+1] {
//		i++
//		count++
//	}
//}

//sorted:= sort.Ints(ar)
//sort.Ints(ar)
//sorted:= sort.Slice(ar, func(i, j int) bool {
//	return ar[i]<ar[j]
//})
func theSockMerchant(n int32, ar []int32) int32 {
	var count int32

	m := make(map[int32]bool)

	for _,v := range ar{
		// Check if the the value was found before
		if m[v] {
			count++
			m[v] = false
		}else {
			// set true for the first item of the pair
			m[v] = true
		}
	}

	return count
}

func theSockMerchant1(n int32, ar []int32) int32 {
	var count int32

	m := make(map[int32]int32)

	for _,v := range ar{
		m[v]++
		if m[v]%2==0 {
			count++
		}
	}

	return count
}
/*
		ar = [1,2, 1,2,1,3,2]
				[1,1,1,2,2,2,3]: 2


*/

func aVeryBigSum(ar []int64) int64 {
	result:=int64(0)
	for i,_:=range ar{
		result+=ar[i]
	}
	return result
}


//type foo []int32
//
//func (ar foo) Len() int {
//	return len(ar)
//}
//
//func (ar foo) Less(i, j int) bool {
//	return ar[i] < ar[j]
//}
//
//func (ar foo) Swap(i, j int) {
//	ar[i], ar[j] = ar[j], ar[i]
//}