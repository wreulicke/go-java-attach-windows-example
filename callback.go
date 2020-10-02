package main

/*
extern void callback(void* ch, char* buf);
void _callback(void* ch, char* buf) {
  callback(ch, buf);
}
extern void closeCallback(void* ch, int r);
void _closeCallback(void* ch, int r) {
  closeCallback(ch, r);
}
*/
import "C"
