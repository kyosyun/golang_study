package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i< len(os.Args) ; i++  {
		fmt.Println("index: " + strconv.Itoa(i) + ", value: "  + os.Args[i])
	}
}
