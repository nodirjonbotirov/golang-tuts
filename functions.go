package main

import (
	"fmt"
)

func main(){
	// Greeting("john")
	//fmt.Println(Sum(1,2))
	sum,multi:=SumAndMultiply(1,2)
	fmt.Println(sum, multi)
}

func Greeting(name string){
	fmt.Printf("Hello mother f*cker: %s\n", name)
}

func Sum(a1, a2 int) int {
	return a1+a2
}
func SumAndMultiply(a,b int)(int, int){
	return a+b, a*b
}