package widget

import "fmt"

func Dall(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
