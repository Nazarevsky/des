package des

import (
	"math"
)

func permutate(block BitArray, permutConstants [][]int) BitArray {
	var permutBlock BitArray
	for _, row := range permutConstants {
		for _, v := range row {
			permutBlock.Append(block.GetBit(v - 1))
		}
	}

	return permutBlock
}

func genRandKeys(key BitArray) []BitArray {
	var keys []BitArray

	keyC := permutate(key, c)
	keyD := permutate(key, d)

	for i := 0; i < RoundAmount; i++ {
		keyC.LShift(keyShift[i])
		keyD.LShift(keyShift[i])

		keyCCopy := keyC

		keys = append(keys, permutate(*keyCCopy.AppendArray(keyD), pc2))
	}
	return keys
}

func bitArrToDecimal(value BitArray) int {
	result := 0
	for _, bit := range value {
		if bit {
			result = (result << 1) | 1
		} else {
			result = result << 1
		}
	}
	return result
}

func decimalToBitArray(decimal int) BitArray {
	var bitArray BitArray

	if decimal == 0 {
		return BitArray{false, false, false, false}
	}

	for decimal > 0 {
		bit := decimal&1 == 1
		bitArray = append([]bool{bit}, bitArray...)
		decimal >>= 1
	}
	bitArray.padToModBE(false, 4)

	return bitArray
}

func sBoxPermutate(block BitArray) BitArray {
	var resBlock BitArray
	for sRound := 0; sRound < 8; sRound++ {
		start := sRound * 6
		sBlock := block[start : start+6]

		var rowBits BitArray
		rowBits.Append(sBlock[0])
		rowBits.Append(sBlock[5])

		row := bitArrToDecimal(rowBits)
		col := bitArrToDecimal(sBlock[1:5])

		resBlock.AppendArray(decimalToBitArray(s[sRound][row][col]))
	}

	resBlock.padToModLE(false, 32)

	return resBlock
}

func getEntropy(data BitArray) float64 {
	zeros, ones := data.Count()
	if ones == 0 || zeros == 0 {
		return 0
	}

	pZeros := float64(zeros) / float64(len(data))
	pOnes := float64(ones) / float64(len(data))

	return -pZeros*math.Log2(pZeros) - pOnes*math.Log2(pOnes)
}
