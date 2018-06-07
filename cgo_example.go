package main

import (
	"fmt"
	"github.com/jwenz723/cgo_example/increment"
	"github.com/jwenz723/cgo_example/stdlib"
	"time"
	"github.com/jwenz723/cgo_example/pdh"
)

func main() {
	i := 3
	fmt.Printf("Increment %d -> %d\n", i, increment.Increment(i))

	stdlib.PrintText()
	stdlib.Seed(time.Now().Nanosecond())
	fmt.Printf("Random number: %d\n", stdlib.Random())

	path := `\Processor(_Total)\% Processor Time`
	ret := pdh.PdhValidatePath(path)
	fmt.Printf("PdhValidatePath(%s) : %x (0 means valid)\n", path, ret)
}