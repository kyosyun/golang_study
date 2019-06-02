package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var g, s, b string
	sc.Scan()
	g = sc.Text()
	sc.Scan()
	s = sc.Text()
	sc.Scan()
	b = sc.Text()
	fmt.Printf("Gold %s\n", g)
	fmt.Printf("Silver %s\n", s)
	fmt.Printf("Bronze %s\n", b)
}
