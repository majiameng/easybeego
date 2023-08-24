package widget

import "fmt"

func Widget(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
