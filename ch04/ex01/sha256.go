package ex01

import (
	"crypto/sha256"
	"gopl.io/ch2/popcount"
)

var popCount [256]byte


func CheckSum(sum1, sum2 [32]byte) int {
	xorBytes := make([]byte, 0, sha256.Size)
	for i := 0; i < sha256.Size; i++ {
		xorBytes = append(xorBytes, sum1[i]^sum2[i])
	}
	return totalPopCount(xorBytes)
}

func totalPopCount(bytes []byte) int {
	total := 0
	for _, b := range bytes {
		total += int(popcount.PopCount(uint64(b)))
	}
	return total
}
