package main

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	for i,_ := range grades {
		if grades[i] >= 38 && ((grades[i]%5 == 3) || (grades[i]%5 == 4)) {
			grades[i] = grades[i] + (5 - (grades[i] % 5))
			//44: 44+ (5-4) =45
			//43: 43+(5-3) =43
		}
	}
	return grades
}
//v:=0
//var f = int64(3)
//f = int64(v)
//v=int(f)