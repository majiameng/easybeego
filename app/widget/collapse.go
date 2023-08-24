package widget

import "fmt"

func Collapse(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
