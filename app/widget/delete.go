package widget

import "fmt"

func Delete(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
