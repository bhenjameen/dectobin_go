package main

import (
	"fmt"
	"strconv"
)

var input int

func main() {
	fmt.Println("DectoBin!")
	fmt.Println("Kindly Input Your Number")

	fmt.Scan(&input)
	n := int64(input)
	y := strconv.FormatInt(n, 2)

	fmt.Println(y)
}
