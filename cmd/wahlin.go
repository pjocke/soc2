package main

import (
	"fmt"

	"github.com/pjocke/wahlin"
)

func main() {
	fmt.Println(wahlin.Encode("philadelphia"))
	fmt.Println(wahlin.Encode("räka"))
	fmt.Println(wahlin.Encode("räkna"))
}
