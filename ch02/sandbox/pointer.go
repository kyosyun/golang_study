package main

import  "fmt"

func main() {
	val := 2
	fmt.Println("val: ", val)
	p := &val
	fmt.Println("pointer: ", p)
	p2 := &p
	fmt.Println("pointer2: ", p2)
	fmt.Println("value of pointer2: ", &p2)
	p3 := &p2
	fmt.Println("pointer3: ", p3)
	fmt.Println("value of pointer3: ", &p3)
}
