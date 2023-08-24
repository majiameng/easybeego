package widget

import "fmt"

func Query(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
