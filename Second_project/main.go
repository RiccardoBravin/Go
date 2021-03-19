package main

import (
	"fmt"
	"lesson_tester/src"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Test one:")
	fn.One()

	fmt.Println("Test two:")
	fn.Two()

	fmt.Println("Test three:")
	fn.Three()

	fmt.Println("Test four:")
	// fmt.Println("Insert two numbers")
	// var a,b float64
	// fmt.Scanln(&a)  //this is how you get input from user
	// fmt.Scanln(&b)
	// a,b = fn.Four(a,b)
	// fmt.Println("The max and min are:", a, "and", b)

	fmt.Println("Test five:")
	bk := fn.Book{
		Title:  "La Commedia",
		Author: "Dante Alighieri",
	}
	fmt.Println("Il libro è", bk)
	fn.Five(&bk, "La Vita Nuova")
	fmt.Println("Il nuovo titolo è", bk)

	fmt.Println("Test six:")
	fn.Six()

	fmt.Println("Test seven:")
	num := rand.Intn(200) - 100
	println(num, "è un numero positivo ->", fn.Seven(num))

	fmt.Println("Test eight:")
	fn.Eight()

	fmt.Println("Test nine:")
	fn.Nine()


}
