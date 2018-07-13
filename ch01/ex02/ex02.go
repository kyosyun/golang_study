package ex02

import (
	"os"
	"fmt"
)

func main() {
	Join(os.Args)
}

func Join(args []string) {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(args[i] + ",")
	}
}