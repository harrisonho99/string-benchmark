package util

import "fmt"

func Brk() {
	fmt.Println("============================")
}

func PrintDetail(i interface{}) {
	fmt.Printf("{type: \"%T\", value: \"%#v\"}\n", i, i)
}
