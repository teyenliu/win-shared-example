//go build -ldflags -v -x -installsuffix cgo -o example.exe .

package main

import (
    "fmt"
)

/*
#cgo windows LDFLAGS: -lwin-shared-example
#cgo windows CFLAGS: -DWINDOWS

#cgo LDFLAGS: -L"./bin/Debug"
#cgo CFLAGS: -I./

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

func main () {
    fmt.Printf("Sum Result: %d", Sum(2, 2))
}