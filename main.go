package main

import (
	"fmt"
)

var name string = "first"

func main() {
	fmt.Println("start")
	fmt.Println(name)
	fmt.Println("end")
}
func init() {
	name = "second"
}
