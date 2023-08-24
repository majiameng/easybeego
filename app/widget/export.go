package widget

import "fmt"

func Export(params ...interface{}) (out string) {
	fmt.Println("---")
	fmt.Println(params)
	return out
}
