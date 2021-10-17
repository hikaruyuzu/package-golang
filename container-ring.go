package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var person *ring.Ring = ring.New(5)
	//
	//person.Value = "Tsukasa Chan"
	//person2 := person.Next()
	//person2.Value = "Kaguya sama"
	//
	//fmt.Println(person.Value)
	//fmt.Println(person.Next().Value)

	for i := 0; i < person.Len(); i++ {
		person.Value = "Data " + strconv.FormatInt(int64(i), 10)
		person.Next()
	}

	person.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
