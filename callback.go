package main

/*
extern void callback(uintptr_t ch, char* buf);
void _callback(uintptr_t ch, char* buf) {
  callback(ch, buf);
}
extern void closeCallback(uintptr_t ch, int r);
void _closeCallback(uintptr_t ch, int r) {
  closeCallback(ch, r);
}
*/
import "C"
