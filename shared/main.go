package main

import (
	"fmt"
	"github.com/andresbott/exploring-CGO/shared/gohello"
)

func main() {
	fmt.Printf("Shared C object code says: \"%s\"\n", gohello.Hello())
}
