package main

import (
	"Lab1_des/des"
	"fmt"
	"math"
)

// Some of the weak keys:
// 1.  - 0101-0101-0101-0101
// 2.  - 1F1F-1F1F-0E0E-0E0E

func getEntropy(data des.BitArray) float64 {
	zeros, ones := data.Count()
	if ones == 0 || zeros == 0 {
		return 0
	}

	pZeros := float64(zeros) / float64(len(data))
	pOnes := float64(ones) / float64(len(data))

	return -pZeros*math.Log2(pZeros) - pOnes*math.Log2(pOnes)
}

func main() {
	//cli.Run()
	//00000000
	//11111111
	//11111111
	//10100001
	//00011111
	//01001010
	//10001000
	//01010101
	//01010101
	// 001001110

	value := des.BitArray{false, true, true, true, false, false, true, false, false}
	println(value.ToString())
	fmt.Printf("entropy: %0.3f\n", getEntropy(value))
}
