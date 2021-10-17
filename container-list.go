package main

import (
	"container/list"
	"fmt"
)

func main() {
	/**
	representasi double linked list -> can't get value with index
	 */

	data := list.New()
	data.PushBack("Tsukasa Chan")
	data.PushBack("Shino Aki")
	data.PushBack("Sagiri")
	data.PushBack("Yukinon")

	data.PushFront("Kaguya Sama")

	// data.Front().Next().Next().Value --> you can access next data with next()
	// data.Back().Prev().Prev().Value --> ----------- before data with prev()
	// access
	//fmt.Println(data.Front().Value)
	//fmt.Println(data.Back().Value)

	// front -> back
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println(data.Remove(data.Back())) // delete

	// back -> front
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}


}
