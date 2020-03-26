package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to DectoBin!")
	fmt.Println("Kindly Input Your Number")

	var input int

	fmt.Scan(&input)
	n := int64(input)
	y := strconv.FormatInt(n, 2)

	fmt.Println(y)
}
