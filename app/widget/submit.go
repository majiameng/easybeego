package widget

import "fmt"

func Submit(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
