package main

import (
	"testing"
)

func Test_judgenum_1(t *testing.T) {
	if judgeNum(28) == "perfect" {
		t.Log("完全数テスト")
	}
	if judgeNum(16) == "neary" {
		t.Log("ほぼ完全数")
	}
	if judgeNum(3) == "neither" {
		t.Log("neither")
	}
}
