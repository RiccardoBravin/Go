package tests

import "fmt"

func noCaps(){
	fmt.Println("this function cant be called outside")
}

func Caps(){
	fmt.Println("this function can be called outside")
}