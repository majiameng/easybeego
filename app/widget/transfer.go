package widget

import "fmt"

func Transfer(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
