package main

import (
	"fmt"
	mod "github.com/saefulrahman/gomodules/v2"
)

func main() {
	hello := mod.SayHello("Saeful Rahman")
	fmt.Println(hello)
}
