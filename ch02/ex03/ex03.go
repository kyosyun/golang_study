package ex03

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	sum := 0
	for i := 0; i < 8 ; i++ {
		sum += int(pc[byte(x>>uint(i*8))])
	}
	return sum
}