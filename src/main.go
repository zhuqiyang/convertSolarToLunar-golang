package main

import (
	"convert"
	"fmt"
)

func main() {
	// convert.Add()
	// a := convert.ConvertSolarToLunar(1998, 4, 14) // 1998-3-18
	// fmt.Println(a)
	// b := convert.ConvertSolarToLunar(1995, 10, 24)

	a := convert.ConvertLunarToSolar(1998, 3, 18, false)
	// fmt.Println(b)
	fmt.Println(a)

}
