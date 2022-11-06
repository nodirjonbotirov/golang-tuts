package main

import (
	"fmt"
)

func main(){
	//var hello string
	//fmt.Println(hello) default empty
	//hello = "Hello"
	//fmt.Println(hello)
	// var ourBool bool
	// fmt.Println(ourBool) // default false
	// ourBool = true
	// fmt.Println(ourBool)
	name := "nodirjon"
	greeting:=fmt.Sprintf("hello %s", name)
	fmt.Printf("Type: %T Value: %v\n", greeting, greeting)

}