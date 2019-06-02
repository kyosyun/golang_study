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
	slice := strings.Split(in, " ")
	from := slice[0]
	to := slice[1]
	fromInt, _ := strconv.Atoi(from)
	toInt, _ := strconv.Atoi(to)
	fmt.Printf("%d", toInt-fromInt)

}
