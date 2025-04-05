package main

import (
	"fmt"

	"piscine"
)

func main() {
	args := [2]string{"DD", "DABC"}
	fmt.Println(piscine.HiddeP(args[0], args[1]))
}
