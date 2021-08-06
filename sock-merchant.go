package main

import (
	"fmt"
)

// concurrency vs parallelism
//rob pike simplicity is complicated
//go doc
//educative.io
// https://invite.slack.golangbridge.org/
//blog
//buy
//ebook
func main() {
	slices1:=[]int32{18,10,2,10,5,11,3,10}
	/*
	{3,2,1,2,1,3,2}
	 */
	//fmt.Println(sockMerchant(7,slices1))
	fmt.Println(sortMethod(slices1))
}
func BubbleSort(array[] int32)[]int32 {
	// first loop iterates from position 0 to last position
	for i:=0; i< len(array)-1; i++ {
		//second loop iterates to second to the last position since last position is sorted already
		for j:=0; j < len(array)-i-1; j++ {
			//swapping occurs if position position[i]> position [i+1]
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

func sortMethod(arr [] int32) []int32 {
	//slices1:=[]int32{3,2,1,5,1,3,2}
	length:= len(arr)
	//sorting with single loop
	for i:=0; i<length-1; i++ {
		if arr[i]>arr[i+1] {
			temp:= arr[i]
			arr[i]=arr[i+1]
			arr[i+1]=temp
			i=-1
		}
	}
	return arr
	//	int length = arr.length;
//	// Sorting using a single loop
//	for (int j = 0; j < length - 1; j++) {
//		// Checking the condition for two
//		// simultaneous elements of the array
//		if (arr[j] > arr[j + 1]) {
//			// Swapping the elements.
//			int temp = arr[j];
//			arr[j] = arr[j + 1];
//			arr[j + 1] = temp;
//
//			// updating the value of j = -1
//			// so after getting updated for j++
//			// in the loop it becomes 0 and
//			// the loop begins from the start.
//			j = -1;
//		}
//	}
//	return arr;
//}

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