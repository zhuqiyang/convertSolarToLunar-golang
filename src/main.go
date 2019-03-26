package main

import (
	"convert"
	"fmt"
)

func main() {
	// convert.Add()
	a := convert.ConvertSolarToLunar(1990, 12, 8)
	fmt.Println(a)
}
