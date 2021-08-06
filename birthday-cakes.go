package main

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	max:=int32(0)
	count:=int32(0)
	for i,_:= range candles{
		if max<candles[i] {
			max=candles[i]
		}
	}
	for i:=0; i< len(candles); i++ {
		if max==candles[i] {
			count++
		}
	}
	return count
}