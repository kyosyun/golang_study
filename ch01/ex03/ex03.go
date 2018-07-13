package ex03

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	Join(os.Args)
}


func Join(args []string) {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(strings.Join(args, ","))
	}
}