package main

import (
	"fmt"
	"regexp"
)

func ValidateEmail(e string) bool  {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}


func main() {
	// referensi
	// https://pkg.go.dev/regexp
	// https://github.com/google/re2/wiki/Syntax

	ValidateEmail("tsukasachan@gmail.com")
	ValidateEmail("tsukasachan@gmailcom")
	ValidateEmail("tsukasachan@gmail.com>>")

	// other example
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])i")

	fmt.Println(regex.MatchString("aki"))
	fmt.Println(regex.MatchString("aiu"))
	fmt.Println(regex.MatchString("aku"))
	fmt.Println(regex.MatchString("aIi"))
	fmt.Println(regex.MatchString("Aai"))

	search := regex.FindAllString("aku ai aui aio uia aiii ani", 5)
	fmt.Println(search)
}
