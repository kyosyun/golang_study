package surface_test

import (
	"golang_study/ch03/surface"
	"math"
	"testing"
)

func TestF(t *testing.T) {
	println(math.MaxFloat64)
	println(surface.F(2.0, 2.0))
}
