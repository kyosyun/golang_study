package comma_test

import (
	"testing"
	"golang_study/ch03/comma"
)

func TestComma(t *testing.T) {
	print(comma.Comma("1235") + "\n")
}