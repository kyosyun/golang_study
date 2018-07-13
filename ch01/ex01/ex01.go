package main

import (
	"fmt"
	"os"
	"ch01"
)

func main() {
	s := ""
	sep := ""
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + sep
		sep = " "
	}
	fmt.Println(s)
}
