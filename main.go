package main

import (
	"fmt"
)
func main() {
	var n int = 21
	fmt.Printf("%T %T %T\n", n, int32(n), string(n))
	fmt.Printf("%t %t %t\n", n, int32(n), string(n))
}

