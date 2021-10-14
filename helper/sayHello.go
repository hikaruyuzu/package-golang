package helper

import "fmt"

func SayHello(name string) string {
	return "Okaerinasai " + name + " , dana Sama"
}

func SetUserData(user, password string) (usern, pass string) {

	var data DataUser
	data.username = user
	data.password = password

	usern = "username : " + data.username
	passwordd := "password: " + data.password

	return usern, passwordd
}

/**
Access modifer, if name variable, function and other start with
Uppercase, -> can access from other package
lowercase, -> can't access from other package and will be error
 */

var version = 1
var Application = "Lear Basic Golang"

func sayGoodBye(name string) { // can't access from other package
	fmt.Println("Hello ", name)
}

func RekursiveFactorial(val int) int {
	if val == 1 {
		return 1
	} else {
		return val * RekursiveFactorial(val-1)
	}
}