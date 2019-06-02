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
	studentNum, _ := strconv.Atoi(strings.Split(in, " ")[0])
	passScore, _ := strconv.Atoi(strings.Split(in, " ")[1])

	for i := 0; i < studentNum; i++ {
		sc.Scan()
		scoreAndAbsentNum := sc.Text()
		score, _ := strconv.Atoi(strings.Split(scoreAndAbsentNum, " ")[0])
		absentNum, _ := strconv.Atoi(strings.Split(scoreAndAbsentNum, " ")[1])
		if calcScore(score, absentNum) >= passScore {
			fmt.Printf("%v\n", i+1)
		}
	}
}

func calcScore(testScore, absentNum int) int {
	score := testScore - absentNum*5
	if score < 0 {
		score = 0
	}
	return score
}
