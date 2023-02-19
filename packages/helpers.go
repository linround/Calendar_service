package packages

import "fmt"

func Helper() {
	fmt.Println("helpers")
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
