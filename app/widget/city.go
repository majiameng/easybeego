package widget

import "fmt"

func City(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
