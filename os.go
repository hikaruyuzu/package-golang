package main

import (
	"fmt"
	"os"
)

func envVariable(key string) string {
	os.Setenv(key, "gopher")
	return os.Getenv(key)
}

func main() {
	value := envVariable("name")
	fmt.Println(value)
	fmt.Printf("os package %s = %s \n","name", value)
	args := os.Args
	fmt.Println(args)
	//fmt.Println(args[1])
	//fmt.Println(args[2])

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname",hostname)
	} else {
		fmt.Println("Error: ", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

	uid := os.Geteuid()
	fmt.Println(uid)

}

