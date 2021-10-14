package main

import (
	"fmt"
	"golang-package/helper"
)

func main() {

	//1. package and import
	result := helper.SayHello("Koi")
	user,pass := helper.SetUserData("Miyuko", "tsukasachan")

	fmt.Println(user, pass)
	fmt.Println(result)

	//2. access modifier
	//helper.sayGoodBye("Miyuki Koishiro") --> will be error because s lowercase(can't access)
	res := helper.RekursiveFactorial(5)
	fmt.Println("Value of recursive ", res) // can access because R UpperCase(can acccess)

	//fmt.Println(helper.version) --> error v lowercase
	fmt.Println("App", helper.Application) // can access because A UpperCase

}
