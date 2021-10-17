package main

import (
	"fmt"
	"strconv"
)

func main() {
	val, err := strconv.ParseBool("vailed")
	if err == nil {
		fmt.Println(val)
	} else {
		fmt.Println("Error: ", err.Error())
	}

	nill, err := strconv.ParseInt("1000000", 10, 32)
	if err == nil {
		fmt.Println(nill)
	} else {
		fmt.Println("Error: ", err.Error())
	}


	intStr := strconv.FormatInt(1000000, 10)
	fmt.Println("string",intStr)

	v := true
	boolStr := strconv.FormatBool(v)
	fmt.Println(boolStr)

	// or direct with
	 att,_ := strconv.Atoi("1000")
	 fmt.Println(att)

	 itt := strconv.Itoa(12)
	 fmt.Println(itt)

}
