package main

import (
	"fmt"
	"github.com/andresbott/exploring-CGO/static/gohello"
)

func main() {
	fmt.Printf("Static C code says: %s\n", gohello.Hello())
	// examples:
	fmt.Printf("- sending and returing a string: \"%s\"\n", gohello.SendString("test string"))
}
