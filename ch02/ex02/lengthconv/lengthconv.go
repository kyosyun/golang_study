package lengthconv

import "fmt"

type Meter float64
type Feet float64

func MtoF(meter Meter) Feet {
	return Feet(meter * 3.2808)
}

func FtoM(f Feet) Meter {
	return Meter(f / 3.2808)
}

func (m Meter) String() string { return fmt.Sprintf("%gm", m)}
func (f Feet) String() string { return fmt.Sprintf("%gf", f)}