package main

import (
	"fmt"
	"golang-package/database"
	// _ "golang-package/database" --> Blank Identifier, use if you just want running init
)

func main() {
	connect := database.GetDatabaseConnection()
	fmt.Println(connect)
}
