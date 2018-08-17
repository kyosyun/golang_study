package ex01_test

import (
	"testing"
	"golang_study/ch04/ex01"
	"crypto/sha256"
)

const (
	sum1 = "12345678901234567890123456789012"
	sum2 = "12345678901234567890123456789013"
	sum3 = "12345678901234567890123456789024"
	sum4 = "32145678901234567890123456789012"
	sum5 = "22345678901234567890123456789012"
)

func TestCheckSum(t *testing.T) {
	println(ex01.CheckSum(sha256.Sum256([]byte(sum1)),sha256.Sum256([]byte(sum1))))
	println(ex01.CheckSum(sha256.Sum256([]byte(sum1)),sha256.Sum256([]byte(sum2))))
	println(ex01.CheckSum(sha256.Sum256([]byte(sum1)),sha256.Sum256([]byte(sum3))))
	println(ex01.CheckSum(sha256.Sum256([]byte(sum1)),sha256.Sum256([]byte(sum4))))
	println(ex01.CheckSum(sha256.Sum256([]byte(sum1)),sha256.Sum256([]byte(sum5))))
}

func TestCheckSum2(t *testing.T) {
	//for _, test = struct {
	//
	//}{}{
	//
	//}
}