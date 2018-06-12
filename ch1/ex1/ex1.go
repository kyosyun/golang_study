package main

import (
	"os"
	"fmt"
)

func main() {
	s := ""
	sep := ""
	for i := 0; i < len(os.Args); i++ {
		s += os.Args[i] + sep
		sep = " "
	}
	fmt.Println(s)
}
