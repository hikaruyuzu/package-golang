package main

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	age int
}

// alias of User

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].age < value[j].age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
func main() {
	// sort
	users := []User {
		{"Tsukasa", 19},
		{"Kaguya", 18},
		{"Yukinon", 17},
	}
	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)


}
