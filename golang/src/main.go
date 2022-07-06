package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abc"
	list := strings.Split(str, "")
	strReverse := ""

	for i := len(list) - 1; i >= 0; i-- {
		strReverse = strReverse + list[i]
	}

	fmt.Print(strReverse)
}
