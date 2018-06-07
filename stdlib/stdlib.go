package stdlib

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import "unsafe"

func PrintText() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

func Random() int {
	return int(C.rand())
}

func Seed(i int) {
	C.srand(C.uint(i))
}