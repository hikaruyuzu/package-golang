package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// date now
	timenow := time.Now()
	fmt.Println(timenow.Year())
	fmt.Println(timenow.Month())
	fmt.Println(timenow.Day())
	fmt.Println(timenow.Hour())
	fmt.Println(timenow.Minute())
	fmt.Println(timenow.Second())
	fmt.Println(timenow.Nanosecond())

	fmt.Println(timenow.Local())

	// set date
	utc := time.Date(2002, 12, 12, 28, 37, 36, 19, time.UTC)
	fmt.Println(utc)

	// convertion
	//layout := time.RFC3339 // or you can write like this , layout := "2006-01-02"
	layout := "2006-01-02"
	convdate, err := time.Parse(layout, "2002-08-24")
	if err == nil {
		fmt.Println(convdate)
	} else {
		fmt.Println("ERROR: ", errors.New("Format yang anda masukan salah"))
	}

	// location

	japan, err := time.LoadLocation("Japan")
	if err == nil {
		fmt.Println("Your location: ", japan)
	} else {
		fmt.Println("ERROR: ", errors.New("Gagal menentukan lokasi"))
	}
}
