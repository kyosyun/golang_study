package popcount_test

import (
"fmt"
"testing"
	"golang_study/ch02/popcount"
)

func TestPopcount(t *testing.T) {
	i := popcount.PopCount(255)
	fmt.Println(i)
}