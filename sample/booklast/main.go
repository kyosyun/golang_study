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
	all, _ := strconv.Atoi(strings.Split(in, " ")[0])
	finish, _ := strconv.Atoi(strings.Split(in, " ")[1])

	fmt.Printf("%v", all-finish)
}
