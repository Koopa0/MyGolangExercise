package main

import (
	"fmt"
	"regexp"
)

func main()  {
	fmt.Println(formatCommas(9527))
	fmt.Println(formatCommas(3345678))
	fmt.Println(formatCommas(-123445))


}
func formatCommas(num int) string {
	str := fmt.Sprintf("%d", num)
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}
