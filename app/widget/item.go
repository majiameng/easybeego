package widget

import "fmt"

func Item(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
