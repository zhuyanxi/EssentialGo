package main

import (
	"fmt"
)

func main() {
	basePath := "/srcache/"
	path := "/srcache/"

	key := path[len(basePath):]
	fmt.Println(key)
}
