package main

import (
	"os"
	"strconv"
	"fmt"
	"golang_study/ch02/ex01"
	"bufio"
	"strings"
	"golang_study/ch02/ex02/lengthconv"
	"golang_study/ch02/ex02/weightconv"
)

func main() {
	args := os.Args[1:]
	if(len(args) == 0) {
		//標準入力を受け付ける
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan(){
			in := stdin.Text()
			if(!printConv(in)){
				fmt.Println("以下の形式で入力してください \n" +
								"W:%f or T:%f or L:%f (W:重さ T:温度, L:長さ) \n" +
								"ex) W:25")
			}
		}
	}
	for _, in := range args {
		printConv(in)
	}
}

func commonParseFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex02: %v\n", err)
		os.Exit(1)
	}
	return f
}

func printConv(in string) bool {
	if(strings.HasPrefix(in, "W:")){
		t := strings.Replace(in, "W:", "", 1)
		printWeight(t)
	} else if (strings.HasPrefix(in, "L:")) {
		t := strings.Replace(in, "L:", "", 1)
		printLength(t)
	} else if (strings.HasPrefix(in, "T:")) {
		t := strings.Replace(in, "T:", "", 1)
		printTemperature(t)
	} else {
		return false
	}
	return true
}

func printTemperature(tmp string) {
	t := commonParseFloat(tmp)
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FtoC(f), c, tempconv.CtoF(c))
}

func printLength(len string) {
	l := commonParseFloat(len)
	meter := lengthconv.Meter(l)
	feet := lengthconv.Feet(l)
	fmt.Printf("%s = %s, %s = %s \n",
		meter, lengthconv.MtoF(meter), feet, lengthconv.FtoM(feet))

}

func printWeight(weight string) {
	w := commonParseFloat(weight)
	k := weightconv.Kilogram(w)
	p := weightconv.Pound(w)
	fmt.Printf("%s = %s, %s = %s \n",
		k, weightconv.KtoP(k), p, weightconv.PtoK(p))


}