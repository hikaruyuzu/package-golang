package main

import (
	"fmt"
	"strings"
)

func main() {

	waifu := "TSUKASA CHAN"
	waifuLow := strings.ToLower(waifu)
	fmt.Println(waifuLow)

	waifuUpp := strings.ToUpper(waifuLow)
	fmt.Println(waifuUpp)

	fmt.Println(strings.Contains("Tsukasa chan", "Tsukasa"))
	fmt.Println(strings.Contains("Tsukasa chan", "tsukasa"))

	fmt.Println(strings.Trim("    Tsukasa Chan Desu Yo                    ", " "))

	fmt.Println(strings.Split("Tsukasa Chan Is my Waifu!", " "))

	fmt.Println(strings.ReplaceAll("my wife is koi chan", "koi", "tsukasa"))
}
