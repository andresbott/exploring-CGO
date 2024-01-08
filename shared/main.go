package main

import (
	"fmt"
	"includeCCode/shared/gohello"
)

func main() {
	fmt.Printf("Shared C object code says: \"%s\"\n", gohello.Hello())
}
