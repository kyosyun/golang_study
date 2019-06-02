package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	in := sc.Text()
	h, _ := strconv.Atoi(strings.Split(in, " ")[0])
	w, _ := strconv.Atoi(strings.Split(in, " ")[1])
	val := make([][]int, h)
	for i := range val {
		val[i] = make([]int, w)
	}
	sc.Scan()
	row := sc.Text()
	val1_1, _ := strconv.Atoi(strings.Split(row, " ")[0])
	val1_2, _ := strconv.Atoi(strings.Split(row, " ")[1])
	sc.Scan()
	row2 := sc.Text()
	val2_1, _ := strconv.Atoi(strings.Split(row2, " ")[0])
	val2_2, _ := strconv.Atoi(strings.Split(row2, " ")[1])

	val[0][0] = val1_1
	val[0][1] = val1_2
	val[1][0] = val2_1
	val[1][1] = val2_2

	for i := 0; i < 2; i++ {
		for j := 2; j < w; j++ {
			val[i][j] = val[i][j-1] + (val[i][1] - val[i][0])
		}
	}

	for i := 2; i < h; i++ {
		for j := 0; j < w; j++ {
			val[i][j] = val[i-1][j] + (val[1][j] - val[0][j])
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j == 0 {
				fmt.Printf("%v", val[i][j])
			} else {
				fmt.Printf(" %v", val[i][j])
			}
		}
		fmt.Printf("\n")
	}
}
