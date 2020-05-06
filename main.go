// Package main provides fmt println hello
package main

import (
	"fmt"

	"github.com/ara-ta3/golang-getting-started/foo"
)

func main() {
	l := foo.Lib()
	fmt.Println(l)
}
