package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		in, _ := strconv.Atoi(sc.Text())
		fmt.Printf("%v\n", judgeNum(in))
	}
}

func judgeNum(num int) string {
	divisorsSum := calcDivisorSum(num)
	if divisorsSum == num {
		return "perfect"
	}

	if int(math.Abs(float64(divisorsSum-num))) == 1 {
		return "nearly"
	}
	return "neither"
}

func calcDivisorSum(num int) int {
	var divisors []int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	var sum = 0
	for i := 0; i < len(divisors); i++ {
		sum += divisors[i]
	}
	return sum
}
