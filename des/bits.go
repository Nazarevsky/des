package des

import (
	"math"
)

type BitArray []bool

func (ba *BitArray) Length() int {
	return len(*ba)
}

func (ba *BitArray) GetBit(pos int) bool {
	return (*ba)[pos]
}

func (ba *BitArray) Append(bit bool) *BitArray {
	*ba = append(*ba, bit)
	return ba
}

func (ba *BitArray) AppendBE(bit bool) *BitArray {
	*ba = append(BitArray{bit}, *ba...)
	return ba
}

func (ba *BitArray) AppendArray(bits []bool) *BitArray {
	*ba = append(*ba, bits...)
	return ba
}

func (ba *BitArray) Clear() {
	*ba = BitArray{}
}

func convBytesToBits(bytes []byte) BitArray {
	var result BitArray
	for _, b := range bytes {
		for i := 7; i >= 0; i-- {
			bit := (b >> uint(i)) & 1
			result.Append(bit == 1)
		}
	}
	return result
}

func (ba BitArray) ToBytes() []byte {
	byteArray := make([]byte, (len(ba)+7)/8)
	for i, b := range ba {
		if b {
			byteArray[i/8] |= 1 << uint(7-i%8)
		}
	}
	return byteArray
}

func (ba BitArray) ToString() string {
	var s string
	for _, b := range ba {
		if b {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}

func (ba *BitArray) padToModLE(value bool, toLen uint) {
	if uint(len(*ba))%toLen == 0 {
		return
	}

	arrLen := uint(len(*ba))
	addAmount := toLen*uint(math.Ceil(float64(arrLen)/float64(toLen))) - arrLen
	var i uint
	for ; i < addAmount; i++ {
		ba.Append(value)
	}
}

func (ba *BitArray) padToModBE(value bool, toLen uint) {
	if uint(len(*ba))%toLen == 0 {
		return
	}

	arrLen := uint(len(*ba))
	addAmount := toLen*uint(math.Ceil(float64(arrLen)/float64(toLen))) - arrLen
	var i uint
	for ; i < addAmount; i++ {
		ba.AppendBE(value)
	}
}

func (ba BitArray) ToBlocks(blockLen uint) []BitArray {
	ba.padToModLE(false, blockLen)

	var blocks []BitArray
	var blockAmount uint
	for ; blockAmount < uint(len(ba))/blockLen; blockAmount++ {
		start := blockAmount * blockLen
		blocks = append(blocks, ba[start:start+blockLen])
	}

	return blocks
}

func (ba *BitArray) LShift(shift int) {
	shiftMod := shift % len(*ba)
	left := (*ba)[:shiftMod]
	right := (*ba)[shiftMod:]

	ba.Clear()
	ba.AppendArray(append(right, left...))
}

func (ba *BitArray) Xor(value BitArray) {
	for i, _ := range *ba {
		(*ba)[i] = (*ba)[i] != value[i]
	}
}

// Count amount of 0 and 1 bits. Returns (amountZeros, amountOnes)
func (ba *BitArray) Count() (int, int) {
	var amountOnes int
	var amountZeros int
	for _, v := range *ba {
		if v {
			amountOnes++
		} else {
			amountZeros++
		}
	}
	return amountZeros, amountOnes
}
