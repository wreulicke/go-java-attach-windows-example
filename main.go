package main

/*
#include "jattach.h"
#cgo LDFLAGS: -L. -ljattach

extern void _callback(uintptr_t ch, char* buf);
extern void _closeCallback(uintptr_t ch, int r);
*/
import "C"
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unsafe"
)

type Channel chan string

//export closeCallback
func closeCallback(p C.uintptr_t, r C.int) {
	fmt.Println("called closeCallback")
	pi := *(*Channel)(unsafe.Pointer(uintptr(p)))
	close(pi)
}

//export callback
func callback(p C.uintptr_t, ch *C.char) {
	fmt.Println("called callback")
	pi := *(*Channel)(unsafe.Pointer(uintptr(p)))
	pi <- C.GoString(ch)
}

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
	ch := make(chan string)
	cCallbacks := C.Callbacks{}
	cCallbacks.callback = C.CallbackFn(C._callback)
	cCallbacks.closeCallback = C.CloseCallbackFn(C._closeCallback)
	pipename := `\\.\pipe\javatool` + os.Args[1]
	go func() {
		r := C.attach(C.uintptr_t(uintptr(unsafe.Pointer(&ch))), cCallbacks, C.int(pid), C.CString(pipename), C.int(len(os.Args)-2), (**C.char)(cArray))
		if r != C.int(0) {
			fmt.Println("attach is failed")
		}
	}()
	for x := range ch {
		fmt.Println(x)
	}
	return nil
}

func main() {
	if err := mainInternal(); err != nil {
		log.Fatalf("%+v", err)
	}
}
