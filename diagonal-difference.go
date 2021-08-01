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
//public static int diagonalDifference(List<List<Integer>> arr) {
//int leftdiagonal = 0, rightdiagonal = 0;
//for(int i = 0, j = arr.get(0).size()-1; i < arr.get(0).size(); i++, j--){
//leftdiagonal = leftdiagonal + arr.get(i).get(i);
//rightdiagonal = rightdiagonal + arr.get(i).get(j);
//}
//return Math.abs(leftdiagonal - rightdiagonal);
//}
