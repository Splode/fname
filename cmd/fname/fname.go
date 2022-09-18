package main

import (
	"fmt"

	"github.com/splode/fname"
)

func main() {
	rng := fname.NewRandomNameGenerator()
	fmt.Println(rng.Generate())
}
