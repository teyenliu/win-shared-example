// hello.go
package main

import "C"

import "fmt"

/* Use cgo to export function for C
   But, it still can be used in Golang
*/
//export SayHelloGo
func SayHelloGo(s *C.char) {
	fmt.Print(C.GoString(s))
}
