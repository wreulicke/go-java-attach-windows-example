package main

/*
#include "jattach.h"
#cgo LDFLAGS: -L. -ljattach
*/
import "C"
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unsafe"
)

func mainInternal() error {
	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return fmt.Errorf("pid is not specified")
	}
	size := len(os.Args) - 2
	cArray := C.malloc(C.size_t(unsafe.Sizeof(uintptr(0))) * C.size_t(size))
	x := (*[3]*C.char)(cArray)
	x[0] = C.CString(os.Args[2])
	x[1] = C.CString(os.Args[3])
	x[2] = C.CString(os.Args[4])
	pipename := `\\.\pipe\javatool` + os.Args[1]
	r := C.attach(C.int(pid), C.CString(pipename), C.int(len(os.Args)-2), (**C.char)(cArray))
	if r != C.int(0) {
		return fmt.Errorf("attach is failed")
	}
	return nil
}

func main() {
	if err := mainInternal(); err != nil {
		log.Fatalf("%+v", err)
	}
}
