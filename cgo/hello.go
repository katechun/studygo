package main

import "C"

//void SayHello(const char* s);
func main() {
	C.SayHello(C.CString("hello,world\n"))
}
