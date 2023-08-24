package widget

import "fmt"

func Edit(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
