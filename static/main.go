package main

import (
	"fmt"
	"includeCCode/static/gohello"
)

func main() {
	fmt.Printf("Static C code says: %s\n", gohello.Hello())
	// examples:
	fmt.Printf("- sending and returing a string: \"%s\"\n", gohello.SendString("test string"))
}
