package widget

import "fmt"

func Addz(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
