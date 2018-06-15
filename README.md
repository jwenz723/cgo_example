# cgo_example
A simple example of how to make use of cgo

# Packages

* [increment](https://github.com/jwenz723/cgo_example/tree/master/increment): demonstrates how a custom C library can be created ([increment.c](https://github.com/jwenz723/cgo_example/blob/master/increment/increment.c), [increment.h](https://github.com/jwenz723/cgo_example/blob/master/increment/increment.h)) and then referenced in Go
* [pdh](https://github.com/jwenz723/cgo_example/tree/master/pdh): demonstrates how a standard windows library can be used in Go
* [stdlib](https://github.com/jwenz723/cgo_example/tree/master/stdlib): demonstrates how a standard library can be used in Go with additional custom C code defined in the C preamble within [stdlib.go](https://github.com/jwenz723/cgo_example/blob/master/stdlib/stdlib.go)
