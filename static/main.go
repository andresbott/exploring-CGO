package main

import (
	"fmt"
	"github.com/andresbott/exploring-CGO/static/gohello"
)

func main() {
	fmt.Printf("Static C code says: %s\n", gohello.Hello())
	// examples:
	fmt.Printf("- different type conversions: \n")
	for _, i := range gohello.TypeConversions() {
		fmt.Printf("|-- %s\n", i)
	}

	fmt.Printf("- sending and returing a string: \"%s\"\n", gohello.SendString("test string"))

	str := gohello.GetStruc()
	fmt.Printf("- reading value from struct: name: \"%s\", numner: \"%d\"\n", str.Name, str.Number)

}
