package ex05

func PopCount(x uint64) int {
	sum := 0
	for x > 0 {
		sum++
		x = x & (x - 1)
	}
	return sum
}