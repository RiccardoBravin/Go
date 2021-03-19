package main

import (
	"fmt"
	"firstprj/src/help" //using the name of the main module (done with -go mod init [name]-) and than the directory
	"firstprj/src/lib_test"

)

func main(){
	fmt.Println("Hello Go!")
	fmt.Println(helpers.HelloWorld("hi"))  //call functions of packages using the package name at the top of the file 
	//tests.noCaps()  functions without caps can't be used outside of scope
	tests.Caps()
}