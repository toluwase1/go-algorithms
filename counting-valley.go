package main

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	//check length of path
	// check for single or continuous D
	altitude:=0
	valley:=0
	for i:=int32(0); i<steps; i++{
		if path[i]== 'U' {
			if altitude==-1 {
				valley++
			}
			altitude++
		}
		if path[i]=='D' {
			altitude--
		}
	}
	return int32(valley)
}
