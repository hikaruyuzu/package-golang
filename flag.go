package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "put your localhostname")
	var user *string = flag.String("name", "root", "put your name")
	var password *string = flag.String("password", "watashinamaewa?", "put your password")
	var number *int = flag.Int("number", 100, "Put your number")

	flag.Parse()
	fmt.Println("host: ",*host)
	fmt.Println("username: ",*user)
	fmt.Println("password: ",*password)
	fmt.Println("number: ", *number)
}
