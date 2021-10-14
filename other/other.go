package other

import "fmt"

func sayHello(name string) {
	// not error, because they are not in the same package
	fmt.Println("Tsukasa Chan")
}
