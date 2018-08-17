package comma_test

import (
	"testing"
	"golang_study/ch03/ex11"
)

func TestComma(t *testing.T) {
	print(comma.CommaWithSinged("1235") + "\n")
	print(comma.CommaWithSinged("+1235") + "\n")
	print(comma.CommaWithSinged("+1235.12345") + "\n")
	print(comma.CommaWithSinged("-1235.12345") + "\n")
	print(comma.CommaWithSinged("+12351234.5") + "\n")
	print(comma.CommaWithSinged("-12351234.5") + "\n")
}
