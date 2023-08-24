package widget

import "fmt"

func Add(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
