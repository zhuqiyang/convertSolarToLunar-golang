# 使用方法
```golang
func main() {
	// 公历转农历
	// a := convert.ConvertSolarToLunar(1998, 4, 14) // 1998-3-18
	// fmt.Println(a)
	// b := convert.ConvertSolarToLunar(1995, 10, 24)
	// fmt.Println(b)
	
	// 农历转公历
	a := convert.ConvertLunarToSolar(1998, 3, 18, false)
	fmt.Println(a)

}
```
