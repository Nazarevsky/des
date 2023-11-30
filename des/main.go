package des

import "errors"

// Encrypt encrypts provided data. Returns bytes of decrypted message, array of entropy, error.
func Encrypt(data []byte, key []byte) ([]byte, [][]float64, error) {
	keyBits := convBytesToBits(key)

	if keyBits.Length() < KeyLength {
		return nil, [][]float64{{}}, errors.New("key length must me equal to 64 bits")
	}

	var res []byte
	var entropys [][]float64

	dataBits := convBytesToBits(data)
	dataBlocks := dataBits.ToBlocks(DataBlockLen)

	rKeys := genRandKeys(keyBits)

	for _, block := range dataBlocks {
		ipBlocks := permutate(block, ip).ToBlocks(DataBlockLen / 2)
		var l BitArray
		var blockEntropys []float64

		for round := 0; round < RoundAmount; round++ {
			l = ipBlocks[1]

			ipBlocks[1] = permutate(ipBlocks[1], e)
			ipBlocks[1].Xor(rKeys[round])
			ipBlocks[1] = sBoxPermutate(ipBlocks[1])
			ipBlocks[1] = permutate(ipBlocks[1], p)
			ipBlocks[1].Xor(ipBlocks[0])

			ipBlocks[0] = l
			blockEntropys = append(blockEntropys, getEntropy(*ipBlocks[0].AppendArray(ipBlocks[1])))
		}
		entropys = append(entropys, blockEntropys)

		res = append(res, permutate(*ipBlocks[1].AppendArray(ipBlocks[0]), ipRev).ToBytes()...)
	}

	return res, entropys, nil
}

// Decrypt decrypts provided data. Returns bytes of decrypted message, array of entropy, error.
func Decrypt(data []byte, key []byte) ([]byte, [][]float64, error) {
	keyBits := convBytesToBits(key)

	if keyBits.Length() < KeyLength {
		return nil, [][]float64{{}}, errors.New("key length must me equal to 64 bits")
	}

	var res []byte
	var entropys [][]float64

	dataBlocks := convBytesToBits(data).ToBlocks(DataBlockLen)
	rKeys := genRandKeys(keyBits)

	for _, block := range dataBlocks {
		ipBlocks := permutate(block, ip).ToBlocks(DataBlockLen / 2)
		var l BitArray
		var blockEntropys []float64

		for round := RoundAmount - 1; round >= 0; round-- {
			l = ipBlocks[1]

			ipBlocks[1] = permutate(ipBlocks[1], e)
			ipBlocks[1].Xor(rKeys[round])
			ipBlocks[1] = sBoxPermutate(ipBlocks[1])
			ipBlocks[1] = permutate(ipBlocks[1], p)
			ipBlocks[1].Xor(ipBlocks[0])

			ipBlocks[0] = l
			blockEntropys = append(blockEntropys, getEntropy(*ipBlocks[0].AppendArray(ipBlocks[1])))
		}
		entropys = append(entropys, blockEntropys)

		res = append(res, permutate(*ipBlocks[1].AppendArray(ipBlocks[0]), ipRev).ToBytes()...)
	}

	i := len(res) - 1
	for ; i >= 0; i-- {
		if res[i] != 0 {
			break
		}
	}

	return res[:i+1], entropys, nil
}

const dataConst = "Glory to Ukraine"

func IsKeyWeak(key []byte) bool {
	dataByte := []byte(dataConst)

	encr1, _, err1 := Encrypt(dataByte, key)
	if err1 != nil {
		panic("encryption must not return error with such a data")
	}

	encr2, _, err2 := Encrypt(encr1, key)
	if err2 != nil {
		panic("encryption must not return error with such a data")
	}

	return string(encr2) == string(dataByte)
}
