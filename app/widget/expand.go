package widget

import "fmt"

func Expand(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
