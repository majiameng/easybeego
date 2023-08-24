package widget

import "fmt"

func Date(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
