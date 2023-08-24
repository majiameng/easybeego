package widget

import "fmt"

func Album(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
