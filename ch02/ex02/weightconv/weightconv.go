package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func KtoP(k Kilogram) Pound {
	return Pound(k * 2.2046)
}

func PtoK(p Pound) Kilogram {
	return Kilogram(p / 2.2046)
}

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k)}
func (p Pound) String() string { return fmt.Sprintf("%glb", p)}