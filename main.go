/*
dynamic link: go build -ldflags -v -x -installsuffix cgo -o example.exe .
static link : go build -ldflags="-extldflags '-static -lstdc++'" -o example_static.exe .
*/

package main

import (
	"fmt"
)

/*
#cgo windows LDFLAGS: -lwin-shared-example
#cgo windows CFLAGS: -DWINDOWS

#cgo LDFLAGS: -L"./bin/Debug"
#cgo CFLAGS: -I./
#cgo LDFLAGS: -lstdc++

#include "example.h"
*/
import "C"

func Sum(a, b int) int {
	return int(C.Sum(C.int(a), C.int(b)))
}

func SumAll(nums ...int) int {
	n := len(nums)
	v := make([]C.int, n)
	for i, _ := range nums {
		v[i] = C.int(nums[i])
	}
	return int(C.SumAll(&v[0], C.int(n)))
}

func main() {
	C.SayHello(C.CString("Hello, World...\n"))
	C.SayHello2(C.CString("Hello 2, World...\n"))
	C.SayHelloGo(C.CString("Hello Go, World...\n"))
	fmt.Printf("Sum Result: %d\n", Sum(2, 2))
}
