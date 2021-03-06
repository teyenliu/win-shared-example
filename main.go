/* the explaination of build cgo on go
dynamic link(on windows): go build -ldflags -v -x -installsuffix cgo -o example.exe .
dynamic link(on linux)  : go build -ldflags -v -x -installsuffix cgo -o example .
static link : go build -ldflags="-extldflags '-static -lstdc++'" -o example_static.exe .
static link : go build -ldflags="-extldflags '-static -lstdc++'" -o example_static .
using xgo   : xgo -targets windows/amd64 github.com/teyenliu/win-shared-example
*/

package main

import (
	"fmt"
)

/*
#cgo windows LDFLAGS: -L./ -lwin-shared-example
#cgo linux LDFLAGS: -L./ -Wl,-rpath=./ -llinux-shared-example
#cgo LDFLAGS: -lstdc++
#cgo windows CFLAGS: -DWINDOWS=1
#cgo linux CFLAGS: -DLINUX=1

#if defined(WINDOWS)
    const char* os = "windows";
#elif defined(LINUX)
    const char* os = "linux";
#else
#    error(unknown os)
#endif

#cgo LDFLAGS: -L${SRCDIR}/bin/Debug
#cgo CFLAGS: -I${SRCDIR}/
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
	fmt.Printf("OS: %s\n", C.GoString(C.os))
}
