package widget

import "fmt"

func Import(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
