package main

import (
	"fmt"

	"github.com/pjocke/soc2/wahlin"
)

func main() {
	fmt.Println(wahlin.Encode("philadelphia"))
	fmt.Println(wahlin.Encode("christian"))
	fmt.Println(wahlin.Encode("afton"))
	fmt.Println(wahlin.Encode("foder"))
	fmt.Println(wahlin.Encode("arg"))
	fmt.Println(wahlin.Encode("get"))
	fmt.Println(wahlin.Encode("häst"))
}
