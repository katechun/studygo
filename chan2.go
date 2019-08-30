package main

var v1 = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好"
	v1 <- true

}

func main() {
	go aGoroutine()
	<-v1
	println(msg)
}
