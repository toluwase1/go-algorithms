package main

import "math"

func diagonalDifference(arr [][]int32) int32 {

	leftDiagonal :=int32(0); rightDiagonal :=int32(0)
	for i:=  range arr {
		//leftDiagonal = leftDiagonal+arr[i][i]
		leftDiagonal +=arr[i][i]
		rightDiagonal +=arr[i][len(arr)-1-i]
	}
	return int32(math.Abs(float64(leftDiagonal - rightDiagonal)))
}

//3-1-0
//3-1-1
//3-1-2
/*
		q 	q	q
		q 	q 	q
		q 	q	q
 */
//
func diagonalDifference2(arr [][]int32) int32 {
	leftdiag:=int32(0)
	rightdiag:=int32(0)
	for i , _ := range arr{
		leftdiag+=arr[i][i]
		rightdiag += arr[i][len(arr)-i-1]
	}
	return int32(int(math.Abs(float64(leftdiag - rightdiag))))
}
